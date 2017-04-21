CryptoPro Signer
================

# Install

``` bash
go get github.com/l-vitaly/cpsigner/...
```

# Usage 

## Sign

``` bash
cpsigner -sha1=4f08119e7dca8db7b8b3fd1b022a6c1593c07ba6 <<< 'sign content here'
```


## Check Sign

``` bash
cpsigner -sha1=4f08119e7dca8db7b8b3fd1b022a6c1593c07ba6 -o=check <<< 'signed content here'
```
