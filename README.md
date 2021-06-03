# canihasvaccine
A small utility to check which year is allowed to vaccinate for Covid19 in the Netherlands

#### get it:
```
go get -u github.com/ranglust/canihasvaccine
```
Or from the [release page](https://github.com/ranglust/canihasvaccine/releases)

### usage
#### Run it!
```shell
prompt> canihasvaccine -y/--year ####
```

#####Example
```shell
promtp> canihasvaccine -y 1981
💉  Yes you can! HOORAY!!! 💉
```

###Config file
you can also place a file in the current directory of the exeutable
named `canihasvaccine.yaml` with the content of the year
####Example
```shell
prompt> cat canihasvaccine.yaml
year: 1981
prompt> canihasvaccine                   
💉  Yes you can! HOORAY!!! 💉
```

## License
[MIT](https://github.com/ranglust/canihasvaccine/blob/main/LICENSE)