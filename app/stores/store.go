package stores

import (
	"context"
	"crypto/tls"
	"database/sql"
	"foodcourt/app/model"
	"io"
	"mime/multipart"

	fiber "github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp"
)

func CreateStore(db *sql.DB) *Store {
	return &Store{
		UserInterface:               NewUserStore(db),
		RestaurantInterface:         NewRestaurantStore(db),
		RestaurantCategoryInterface: NewRestaurantCategoryStore(db),
		RolesInterface:              NewRolesStore(db),
	}
}

type Store struct {
	model.UserInterface
	model.RestaurantCategoryInterface
	model.RestaurantInterface
	model.RolesInterface
}

// Accepts implements fiber.Ctx.
func (*Store) Accepts(offers ...string) string {
	panic("unimplemented")
}

// AcceptsCharsets implements fiber.Ctx.
func (*Store) AcceptsCharsets(offers ...string) string {
	panic("unimplemented")
}

// AcceptsEncodings implements fiber.Ctx.
func (*Store) AcceptsEncodings(offers ...string) string {
	panic("unimplemented")
}

// AcceptsLanguages implements fiber.Ctx.
func (*Store) AcceptsLanguages(offers ...string) string {
	panic("unimplemented")
}

// App implements fiber.Ctx.
func (*Store) App() *fiber.App {
	panic("unimplemented")
}

// Append implements fiber.Ctx.
func (*Store) Append(field string, values ...string) {
	panic("unimplemented")
}

// Attachment implements fiber.Ctx.
func (*Store) Attachment(filename ...string) {
	panic("unimplemented")
}

// AutoFormat implements fiber.Ctx.
func (*Store) AutoFormat(body any) error {
	panic("unimplemented")
}

// BaseURL implements fiber.Ctx.
func (*Store) BaseURL() string {
	panic("unimplemented")
}

// Bind implements fiber.Ctx.
func (*Store) Bind() *fiber.Bind {
	panic("unimplemented")
}

// BindVars implements fiber.Ctx.
func (*Store) BindVars(vars fiber.Map) error {
	panic("unimplemented")
}

// Body implements fiber.Ctx.
func (*Store) Body() []byte {
	panic("unimplemented")
}

// ClearCookie implements fiber.Ctx.
func (*Store) ClearCookie(key ...string) {
	panic("unimplemented")
}

// ClientHelloInfo implements fiber.Ctx.
func (*Store) ClientHelloInfo() *tls.ClientHelloInfo {
	panic("unimplemented")
}

// Context implements fiber.Ctx.
func (*Store) Context() *fasthttp.RequestCtx {
	panic("unimplemented")
}

// Cookie implements fiber.Ctx.
func (*Store) Cookie(cookie *fiber.Cookie) {
	panic("unimplemented")
}

// Cookies implements fiber.Ctx.
func (*Store) Cookies(key string, defaultValue ...string) string {
	panic("unimplemented")
}

// Download implements fiber.Ctx.
func (*Store) Download(file string, filename ...string) error {
	panic("unimplemented")
}

// FormFile implements fiber.Ctx.
func (*Store) FormFile(key string) (*multipart.FileHeader, error) {
	panic("unimplemented")
}

// FormValue implements fiber.Ctx.
func (*Store) FormValue(key string, defaultValue ...string) string {
	panic("unimplemented")
}

// Format implements fiber.Ctx.
func (*Store) Format(handlers ...fiber.ResFmt) error {
	panic("unimplemented")
}

// Fresh implements fiber.Ctx.
func (*Store) Fresh() bool {
	panic("unimplemented")
}

// Get implements fiber.Ctx.
func (*Store) Get(key string, defaultValue ...string) string {
	panic("unimplemented")
}

// GetRespHeader implements fiber.Ctx.
func (*Store) GetRespHeader(key string, defaultValue ...string) string {
	panic("unimplemented")
}

// GetRouteURL implements fiber.Ctx.
func (*Store) GetRouteURL(routeName string, params fiber.Map) (string, error) {
	panic("unimplemented")
}

// Host implements fiber.Ctx.
func (*Store) Host() string {
	panic("unimplemented")
}

// Hostname implements fiber.Ctx.
func (*Store) Hostname() string {
	panic("unimplemented")
}

// IP implements fiber.Ctx.
func (*Store) IP() string {
	panic("unimplemented")
}

// IPs implements fiber.Ctx.
func (*Store) IPs() (ips []string) {
	panic("unimplemented")
}

// Is implements fiber.Ctx.
func (*Store) Is(extension string) bool {
	panic("unimplemented")
}

// IsFromLocal implements fiber.Ctx.
func (*Store) IsFromLocal() bool {
	panic("unimplemented")
}

// IsProxyTrusted implements fiber.Ctx.
func (*Store) IsProxyTrusted() bool {
	panic("unimplemented")
}

// JSON implements fiber.Ctx.
func (*Store) JSON(data any, ctype ...string) error {
	panic("unimplemented")
}

// JSONP implements fiber.Ctx.
func (*Store) JSONP(data any, callback ...string) error {
	panic("unimplemented")
}

// Links implements fiber.Ctx.
func (*Store) Links(link ...string) {
	panic("unimplemented")
}

// Locals implements fiber.Ctx.
func (*Store) Locals(key any, value ...any) (val any) {
	panic("unimplemented")
}

// Location implements fiber.Ctx.
func (*Store) Location(path string) {
	panic("unimplemented")
}

// Method implements fiber.Ctx.
func (*Store) Method(override ...string) string {
	panic("unimplemented")
}

// MultipartForm implements fiber.Ctx.
func (*Store) MultipartForm() (*multipart.Form, error) {
	panic("unimplemented")
}

// Next implements fiber.Ctx.
func (*Store) Next() (err error) {
	panic("unimplemented")
}

// OriginalURL implements fiber.Ctx.
func (*Store) OriginalURL() string {
	panic("unimplemented")
}

// Params implements fiber.Ctx.
func (*Store) Params(key string, defaultValue ...string) string {
	panic("unimplemented")
}

// ParamsInt implements fiber.Ctx.
func (*Store) ParamsInt(key string, defaultValue ...int) (int, error) {
	panic("unimplemented")
}

// Path implements fiber.Ctx.
func (*Store) Path(override ...string) string {
	panic("unimplemented")
}

// Port implements fiber.Ctx.
func (*Store) Port() string {
	panic("unimplemented")
}

// Protocol implements fiber.Ctx.
func (*Store) Protocol() string {
	panic("unimplemented")
}

// Queries implements fiber.Ctx.
func (*Store) Queries() map[string]string {
	panic("unimplemented")
}

// Query implements fiber.Ctx.
func (*Store) Query(key string, defaultValue ...string) string {
	panic("unimplemented")
}

// Range implements fiber.Ctx.
func (*Store) Range(size int) (rangeData fiber.Range, err error) {
	panic("unimplemented")
}

// Redirect implements fiber.Ctx.
func (*Store) Redirect() *fiber.Redirect {
	panic("unimplemented")
}

// Render implements fiber.Ctx.
func (*Store) Render(name string, bind fiber.Map, layouts ...string) error {
	panic("unimplemented")
}

// Request implements fiber.Ctx.
func (*Store) Request() *fasthttp.Request {
	panic("unimplemented")
}

// Reset implements fiber.Ctx.
func (*Store) Reset(fctx *fasthttp.RequestCtx) {
	panic("unimplemented")
}

// Response implements fiber.Ctx.
func (*Store) Response() *fasthttp.Response {
	panic("unimplemented")
}

// RestartRouting implements fiber.Ctx.
func (*Store) RestartRouting() error {
	panic("unimplemented")
}

// Route implements fiber.Ctx.
func (*Store) Route() *fiber.Route {
	panic("unimplemented")
}

// SaveFile implements fiber.Ctx.
func (*Store) SaveFile(fileheader *multipart.FileHeader, path string) error {
	panic("unimplemented")
}

// SaveFileToStorage implements fiber.Ctx.
func (*Store) SaveFileToStorage(fileheader *multipart.FileHeader, path string, storage fiber.Storage) error {
	panic("unimplemented")
}

// Scheme implements fiber.Ctx.
func (*Store) Scheme() string {
	panic("unimplemented")
}

// Secure implements fiber.Ctx.
func (*Store) Secure() bool {
	panic("unimplemented")
}

// Send implements fiber.Ctx.
func (*Store) Send(body []byte) error {
	panic("unimplemented")
}

// SendFile implements fiber.Ctx.
func (*Store) SendFile(file string, compress ...bool) error {
	panic("unimplemented")
}

// SendStatus implements fiber.Ctx.
func (*Store) SendStatus(status int) error {
	panic("unimplemented")
}

// SendStream implements fiber.Ctx.
func (*Store) SendStream(stream io.Reader, size ...int) error {
	panic("unimplemented")
}

// SendString implements fiber.Ctx.
func (*Store) SendString(body string) error {
	panic("unimplemented")
}

// Set implements fiber.Ctx.
func (*Store) Set(key string, val string) {
	panic("unimplemented")
}

// SetUserContext implements fiber.Ctx.
func (*Store) SetUserContext(ctx context.Context) {
	panic("unimplemented")
}

// Stale implements fiber.Ctx.
func (*Store) Stale() bool {
	panic("unimplemented")
}

// Status implements fiber.Ctx.
func (*Store) Status(status int) fiber.Ctx {
	panic("unimplemented")
}

// String implements fiber.Ctx.
func (*Store) String() string {
	panic("unimplemented")
}

// Subdomains implements fiber.Ctx.
func (*Store) Subdomains(offset ...int) []string {
	panic("unimplemented")
}

// Type implements fiber.Ctx.
func (*Store) Type(extension string, charset ...string) fiber.Ctx {
	panic("unimplemented")
}

// UserContext implements fiber.Ctx.
func (*Store) UserContext() context.Context {
	panic("unimplemented")
}

// Vary implements fiber.Ctx.
func (*Store) Vary(fields ...string) {
	panic("unimplemented")
}

// Write implements fiber.Ctx.
func (*Store) Write(p []byte) (int, error) {
	panic("unimplemented")
}

// WriteString implements fiber.Ctx.
func (*Store) WriteString(s string) (int, error) {
	panic("unimplemented")
}

// Writef implements fiber.Ctx.
func (*Store) Writef(f string, a ...any) (int, error) {
	panic("unimplemented")
}

// XHR implements fiber.Ctx.
func (*Store) XHR() bool {
	panic("unimplemented")
}

// XML implements fiber.Ctx.
func (*Store) XML(data any) error {
	panic("unimplemented")
}

// release implements fiber.Ctx.
func (*Store) release() {
	panic("unimplemented")
}

// setReq implements fiber.Ctx.
func (*Store) setReq(fctx *fasthttp.RequestCtx) {
	panic("unimplemented")
}
