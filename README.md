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

You need to design a Vending Machine which
```
Accepts coins of 1,5,10,25 Cents i.e. penny, nickel, dime, and quarter.
Allow user to select products Coke(25), Pepsi(35), Soda(45)
Allow user to take refund by canceling the request.
Return the selected product and remaining change if any
Allow reset operation for vending machine supplier.

Reference
Read more: https://javarevisited.blogspot.com/2016/06/design-vending-machine-in-java.html#ixzz6tk72THaA
```
