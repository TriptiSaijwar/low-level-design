# low-level-design


Github Configurations for go get to work (Through ssh)

```
  $ git config --global url.git@github.com:.insteadOf https://github.com/
  $ cat ~/.gitconfig
  [url "git@github.com:"]
      insteadOf = https://github.com/
  $ go get github.com/private/repo
```

For Go modules to work (with Go 1.11 or newer), you'll also need to set the GOPRIVATE variable, to avoid using the public servers to fetch the code:
 
```
  export GOPRIVATE=github.com/private/repo
```
