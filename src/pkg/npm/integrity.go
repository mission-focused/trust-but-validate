package npmintegrity

func Validate(mod string, tag string) error {

	// Here we need to replicate:
	// resolved: https://registry.npmjs.org/autoprefixer/-/autoprefixer-10.4.19.tgz
	// version: 10.4.19
	// curl -LO https://registry.npmjs.org/autoprefixer/-/autoprefixer-10.4.19.tgz

	// cat ./electron-to-chromium-1.4.698.tgz | openssl dgst -sha512 -binary | openssl base64 -A
	return nil
}
