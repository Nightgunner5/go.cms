#!/bin/bash

go get -v github.com/nsf/bin2go

export PATH=$PATH:$GOPATH/bin

wget https://ajax.googleapis.com/ajax/libs/jquery/1/jquery.min.js
wget http://twitter.github.com/bootstrap/assets/bootstrap.zip
unzip bootstrap.zip

echo 'body{padding:60px 0 20px}' > bootstrap/css/bootstrap-extra.min.css

cat bootstrap/css/bootstrap.min.css bootstrap/css/bootstrap-extra.min.css bootstrap/css/bootstrap-responsive.min.css | bin2go -out resource_bootstrap.css.go -pkg http BootstrapCSS
echo 'CSS written'

cat jquery.min.js bootstrap/js/bootstrap.min.js | bin2go -out resource_bootstrap.js.go -pkg http BootstrapJS
echo 'JavaScript written'

bin2go -in bootstrap/img/glyphicons-halflings-white.png -pkg http -out resource_bootstrap.png_white.go GlyphIconsWhite
bin2go -in bootstrap/img/glyphicons-halflings.png -pkg http -out resource_bootstrap.png.go GlyphIcons
echo 'images written'

# Clean up
rm -r bootstrap bootstrap.zip jquery.min.js
