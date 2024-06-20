package translation

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/leonelquinteros/gotext"
	"golang.org/x/text/language"
)

type contextKey string

var userLocalesKey contextKey = "userLocales"

// TranslatorInterface should be implemented by Translator struct
type TranslatorInterface interface {
	Get(key string, locales []string) string
	GetForContext(ctx context.Context, key string) string
}

// Translator handles translations
type Translator struct {
	defaultLocale string
	locales       map[string]*gotext.Locale
}

// NewTranslator instanciates a new Translator struct.
// Locales should be passed by priority order.
// First locale will be the default one.
// First locale for a specific language will be the default for
// this language.
func NewTranslator(path string, locales []string) *Translator {
	localeObjs := map[string]*gotext.Locale{}
	defaultTag, _ := language.Parse(locales[0])
	for _, locale := range locales {
		tag, err := language.Parse(locale)
		if err != nil {
			log.Printf("Unknown locale '%s'", locale)
			continue
		}
		// Setting default locale for a specific language
		language, _ := tag.Base()
		if _, ok := localeObjs[language.String()]; !ok {
			localeObjs[language.String()] = gotext.NewLocale(path, locale)
			go localeObjs[language.String()].AddDomain("lang")
		}
		// Setting locale
		localeObjs[tag.String()] = gotext.NewLocale(path, locale)
		// AddDomain loads and parse the po file, hence the goroutine
		// to keep api-template startup fast
		go localeObjs[tag.String()].AddDomain("lang")
	}
	return &Translator{
		defaultLocale: defaultTag.String(),
		locales:       localeObjs,
	}
}

// Get gets the translation for a given key and locales
// ordered by preference. The first match returns.
// Fallback to the default locale if the ones given
// does not exist
func (t *Translator) Get(key string, locales []string) string {
	for _, locale := range locales {
		tag, err := language.Parse(locale)
		if err != nil {
			continue
		}
		value, ok := t.locales[tag.String()]
		if ok {
			return value.Get(key)
		}
		language, _ := tag.Base()
		value, ok = t.locales[language.String()]
		if ok {
			return value.Get(key)
		}
	}
	return t.locales[t.defaultLocale].Get(key)
}

// GetForContext gets the translation for a given key and context.
// Same logic as Get().
func (t *Translator) GetForContext(ctx context.Context, key string) string {
	userLocales := GetUserLocalesForContext(ctx)
	return t.Get(key, userLocales)
}

// ParseAcceptLanguageMiddleware parses the Accept-Language http header
func ParseAcceptLanguageMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tags, _, err := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
		if err != nil {
			next.ServeHTTP(w, r)
		}
		locales := []string{}
		for _, tag := range tags {
			locales = append(locales, tag.String())
		}
		ctx := context.WithValue(r.Context(), userLocalesKey, locales)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserLocalesForContext finds the user locales from the context.
func GetUserLocalesForContext(ctx context.Context) []string {
	locales, _ := ctx.Value(userLocalesKey).([]string)
	return locales
}

// LocaleForContext finds the locale from the context.
func LocaleForContext(ctx context.Context) string {
	locale, _ := ctx.Value(userLocalesKey).([]string)
	if locale != nil {
		return strings.ToLower(locale[0])
	}
	return "en"
}
