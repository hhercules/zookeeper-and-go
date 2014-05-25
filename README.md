## Zookeeper and Go

Source code and Vagrant configurations for getting started with
Zookeeper and Go.

Please see the [Getting Started with Zookeeper and Go](https://mmcgrana.github.io/.../getting-started-with-zookeeper-and-go.html)
blog post for details on this code and using Zookeeper with Go.

The basic flow is to install a recent version of Vagrant with
Virtualbox and then:

```console
$ vagrant up
$ vagrant ssh go
(go) ~ $ go get github.com/samuel/go-zookeeper/zk
(go) ~ $ go get github.com/mmcgrana/zk
(go) ~ $ go run /vagrant/ex-ping.go
```

### Todo

failure simulation: no zookeper running
  -> error on connect
failure simulation: kill minority of servers
  -> everything works, some weird logging?
failure simulation: kill majority of servers
  -> can read, can't write?
failure simulation: restart cluster
  -> how does it even work
zookeeper acls?
zk library documentation pull requests
blog draft
blog peer review
blog figure out go highlighting
blog publication
blog marketing