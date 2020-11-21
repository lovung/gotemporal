# Golang package to implement the Temporal Model with ORM

## Structure
```
├── LICENSE
├── README.md
├── bitemporal
│   ├── bi.go
│   ├── bi_gormer.go
│   └── collection.go
├── errors.go
├── example
├── go.mod
├── go.sum
├── gorm_callbacks.go
├── gotemporal.go
├── model.go
└── unitemporal
    ├── sys_uni_gormer.go
    └── uni.go
```
## Models
- [ ] Bi-temporal
- [ ] Uni-temporal with system time
- [ ] Uni-temporal with valid time

## ORM
- [ ] GORM
- [ ] XORM

## Referrence
- [A Deep Dive into Bitemporal](https://www.marklogic.com/blog/bitemporal/)
- [What is bitemporal and why should the enterprise care?](https://www.networkworld.com/article/3186634/what-is-bitemporal-and-why-should-the-enterprise-care.html)
