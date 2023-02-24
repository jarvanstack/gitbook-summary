# Install gitbook-summary
if ! command -v gitbook-summary &> /dev/null
then
    # If have golang installed, use go install
    if command -v go &> /dev/null
    then
        go install github.com/dengjiawen8955/gitbook-summary@latest
    else
        # Download binary from github.com/dengjiawen8955/gitbook-summary/bin
        set -e
        wget https://github.com/dengjiawen8955/gitbook-summary/raw/master/bin/gitbook-summary
        sudo mv gitbook-summary /usr/local/bin
        sudo chmod +x /usr/local/bin/gitbook-summary
    fi
    printf "gitbook-summary installed successfully\n"
else 
    printf "gitbook-summary already installed\n"    
fi
gitbook-summary -v
