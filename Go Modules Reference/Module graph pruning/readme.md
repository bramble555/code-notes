go 1.17 引入 懒加载模块

example:

我这个项目 引用 a 和 b 依赖项,a 和 b  都依赖 c,但是 go.mod 里面 没有 

require  c // indirect

不报错,这就是懒加载(前提是 a 和 b 依赖的 c 版本不冲突)