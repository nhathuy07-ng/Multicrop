# Note that pkgconfig has been patched to remove libintl dependency

export CGO_CFLAGS="-D_GNU_SOURCE -Wno-error=implicit-function-declaration -Wno-error=unused-variable \
-I/home/huyn/.asdf/installs/golang/1.26.5/packages/pkg/mod/github.com/davidbyttow/govips/v2@v2.20.0/vips/ \
-I/home/huyn/Downloads/vips-dev-8.19/include/ \
-I/home/huyn/Downloads/vips-dev-8.19/include/glib-2.0/ \
-I/home/huyn/Downloads/vips-dev-8.19/include/glib-2.0/gio/ \
-I/home/huyn/Downloads/vips-dev-8.19/include/glib-2.0/glib/ \
-I/home/huyn/Downloads/vips-dev-8.19/include/glib-2.0/glib/deprecated/ \
-I/home/huyn/Downloads/vips-dev-8.19/lib/glib-2.0/include/ \
-I/home/huyn/Downloads/vips-dev-8.19/lib/"

# Instruct mingw32-gcc's linker step where to find the Windows version of the libraries
export CGO_LDFLAGS="-L/home/huyn/Downloads/vips-dev-8.19/lib"

export PKG_CONFIG_PATH="/home/huyn/Downloads/vips-dev-8.19/lib/pkgconfig/" 

export PKG_CONFIG_ALLOW_CROSS=1 
export CC="x86_64-w64-mingw32-gcc"
export CXX="x86_64-w64-mingw32-g++"
export CGO_ENABLED=1 
export GOOS="windows" 
export GOARCH="amd64"
wails build