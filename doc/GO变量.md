# GO 变量
> 变量: 使用一个名字来绑定一块内存地址, 变量值可以修改
> 关键字 **`var`**
## 显式声明
### 完整声明
`var varName dataType [ = value ]`
- var 关键字
- varName 变量名
- dataType 数据类型
- value 变量初始值

### 短声明
`varName := value`
- `:=` 声明只能在函数内出现
- varName 的数据类型 go 编译器会自行推论

---
# GO 常量
> 常量: 使用一个名字来绑定一块内存地址, 常量值不能修改
> 关键字 **`const``**
```go
const (
 varName = true
)

```