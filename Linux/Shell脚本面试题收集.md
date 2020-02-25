# Shell 脚本面试题

- [Shell 脚本面试题](#shell-%e8%84%9a%e6%9c%ac%e9%9d%a2%e8%af%95%e9%a2%98)
  - [shell 变量](#shell-%e5%8f%98%e9%87%8f)
    - [只读变量(readonly)](#%e5%8f%aa%e8%af%bb%e5%8f%98%e9%87%8freadonly)
    - [删除变量(unset)](#%e5%88%a0%e9%99%a4%e5%8f%98%e9%87%8funset)
    - [变量类型](#%e5%8f%98%e9%87%8f%e7%b1%bb%e5%9e%8b)
    - [shell 字符串](#shell-%e5%ad%97%e7%ac%a6%e4%b8%b2)
    - [拼接字符串](#%e6%8b%bc%e6%8e%a5%e5%ad%97%e7%ac%a6%e4%b8%b2)
    - [获取字符串长度](#%e8%8e%b7%e5%8f%96%e5%ad%97%e7%ac%a6%e4%b8%b2%e9%95%bf%e5%ba%a6)
    - [提取子字符串](#%e6%8f%90%e5%8f%96%e5%ad%90%e5%ad%97%e7%ac%a6%e4%b8%b2)
    - [查找子字符串](#%e6%9f%a5%e6%89%be%e5%ad%90%e5%ad%97%e7%ac%a6%e4%b8%b2)
    - [Shell 数组](#shell-%e6%95%b0%e7%bb%84)
    - [Shell 注释](#shell-%e6%b3%a8%e9%87%8a)
  - [Shell 传递参数](#shell-%e4%bc%a0%e9%80%92%e5%8f%82%e6%95%b0)
  - [Shell 数组](#shell-%e6%95%b0%e7%bb%84-1)
  - [Shell 基本运算符](#shell-%e5%9f%ba%e6%9c%ac%e8%bf%90%e7%ae%97%e7%ac%a6)
    - [算术运算符](#%e7%ae%97%e6%9c%af%e8%bf%90%e7%ae%97%e7%ac%a6)
    - [关系运算符](#%e5%85%b3%e7%b3%bb%e8%bf%90%e7%ae%97%e7%ac%a6)
    - [布尔运算符](#%e5%b8%83%e5%b0%94%e8%bf%90%e7%ae%97%e7%ac%a6)
    - [逻辑运算符](#%e9%80%bb%e8%be%91%e8%bf%90%e7%ae%97%e7%ac%a6)
    - [字符串运算符](#%e5%ad%97%e7%ac%a6%e4%b8%b2%e8%bf%90%e7%ae%97%e7%ac%a6)
    - [文件测试运算符](#%e6%96%87%e4%bb%b6%e6%b5%8b%e8%af%95%e8%bf%90%e7%ae%97%e7%ac%a6)
  - [Shell 流程控制](#shell-%e6%b5%81%e7%a8%8b%e6%8e%a7%e5%88%b6)
    - [if else](#if-else)
    - [for 循环](#for-%e5%be%aa%e7%8e%af)
    - [while 语句](#while-%e8%af%ad%e5%8f%a5)
    - [until 循环](#until-%e5%be%aa%e7%8e%af)
    - [case ... esac](#case--esac)
    - [跳出循环](#%e8%b7%b3%e5%87%ba%e5%be%aa%e7%8e%af)
  - [Shell 函数](#shell-%e5%87%bd%e6%95%b0)
  - [Shell 输入/输出重定向](#shell-%e8%be%93%e5%85%a5%e8%be%93%e5%87%ba%e9%87%8d%e5%ae%9a%e5%90%91)
  - [Shell 文件包含](#shell-%e6%96%87%e4%bb%b6%e5%8c%85%e5%90%ab)
  - [Shell 中各种括号的作用(),(()),\[\],\[\[\]\],{}](#shell-%e4%b8%ad%e5%90%84%e7%a7%8d%e6%8b%ac%e5%8f%b7%e7%9a%84%e4%bd%9c%e7%94%a8)
    - [小括号](#%e5%b0%8f%e6%8b%ac%e5%8f%b7)
      - [单小括号](#%e5%8d%95%e5%b0%8f%e6%8b%ac%e5%8f%b7)
      - [双小括号](#%e5%8f%8c%e5%b0%8f%e6%8b%ac%e5%8f%b7)
    - [中括号](#%e4%b8%ad%e6%8b%ac%e5%8f%b7)
      - [单中括号](#%e5%8d%95%e4%b8%ad%e6%8b%ac%e5%8f%b7)
      - [双中括号](#%e5%8f%8c%e4%b8%ad%e6%8b%ac%e5%8f%b7)
    - [大括号](#%e5%a4%a7%e6%8b%ac%e5%8f%b7)
      - [常规用法](#%e5%b8%b8%e8%a7%84%e7%94%a8%e6%b3%95)
      - [几种特殊的替换结构](#%e5%87%a0%e7%a7%8d%e7%89%b9%e6%ae%8a%e7%9a%84%e6%9b%bf%e6%8d%a2%e7%bb%93%e6%9e%84)
      - [四种模式匹配替换结构](#%e5%9b%9b%e7%a7%8d%e6%a8%a1%e5%bc%8f%e5%8c%b9%e9%85%8d%e6%9b%bf%e6%8d%a2%e7%bb%93%e6%9e%84)
      - [字符串提取和替换](#%e5%ad%97%e7%ac%a6%e4%b8%b2%e6%8f%90%e5%8f%96%e5%92%8c%e6%9b%bf%e6%8d%a2)
    - [符号\$后的括号](#%e7%ac%a6%e5%8f%b7%e5%90%8e%e7%9a%84%e6%8b%ac%e5%8f%b7)
    - [使用](#%e4%bd%bf%e7%94%a8)
  - [shell 面试题](#shell-%e9%9d%a2%e8%af%95%e9%a2%98)
  - [参考](#%e5%8f%82%e8%80%83)

## shell 变量

### 只读变量(readonly)

```sh
#!/bin/bash
myUrl="http://www.google.com"
readonly myUrl
myUrl="http://www.runoob.com"

# output
# /bin/sh: NAME: This variable is read only.
```

### 删除变量(unset)

```sh
#!/bin/sh
myUrl="http://www.runoob.com"
unset myUrl
echo $myUrl
```

### 变量类型

- **局部变量**：在脚本或命令中定义，仅在当前 shell 实例中有效，其他 shell 启动的程序不能访问局部变量。
- **环境变量**：所有程序，包括 shell 启动程序，都能访问环境变量，有些程序需要环境变量来保证其正常运行。必要的时候 shell 脚本也可以定义环境变量。
- **shell 变量**： 是由 shell 程序设置的特殊变量。shell 变量中有一部分是环境变量，有一部分是局部变量，这些变量保证了 shell 的正常运行

### shell 字符串

**单引号**

```sh
str='this is a string'
```

- 单引号里的任何字符都会原样输出，单引号字符串中的变量是无效的；
- 单引号字串中不能出现单独一个的单引号（对单引号使用转义符后也不行），但可成对出现，作为字符串拼接使用。

**双引号**

```sh
your_name='runoob'
str="Hello, I know you are \"$your_name\"! \n"
echo -e $str

# output
# Hello, I know you are "runoob"!
```

- 双引号里可以有变量
- 双引号里可以出现转义字符

### 拼接字符串

```sh
your_name="runoob"
# 使用双引号拼接
greeting="hello, "$your_name" !"
greeting_1="hello, ${your_name} !"
echo $greeting  $greeting_1
# 使用单引号拼接
greeting_2='hello, '$your_name' !'
greeting_3='hello, ${your_name} !'
echo $greeting_2  $greeting_3

# hello, runoob ! hello, runoob !
# hello, runoob ! hello, ${your_name} !
```

### 获取字符串长度

```sh
string="abcd"
echo ${#string} #输出 4
```

### 提取子字符串

```sh
string="runoob is a great site"
echo ${string:1:4} # 输出 unoo
```

### 查找子字符串

```sh
string="runoob is a great site"
echo `expr index "$string" io`  # 输出 4
```

### Shell 数组

bash 支持一维数组（不支持多维数组），并且没有限定数组的大小。

类似于 C 语言，数组元素的下标由 0 开始编号。获取数组中的元素要利用下标，下标可以是整数或算术表达式，其值应大于或等于 0。

**定义数组**

数组名=(值 1 值 2 ... 值 n)

```sh
array_name=(value0 value1 value2 value3)

# or
array_name=(
value0
value1
value2
value3
)

# 还可以单独定义数组的各个分量：
# 可以不使用连续的下标，而且下标的范围没有限制。
array_name[0]=value0
array_name[1]=value1
array_name[n]=valuen
```

**读取数组**

`${数组名[下标]}`

```sh
valuen=${array_name[n]}

使用 @ 符号可以获取数组中的所有元素，例如
echo ${array_name[@]}
```

**获取数组的长度**

```sh
# 取得数组元素的个数
length=${#array_name[@]}
# 或者
length=${#array_name[*]}
# 取得数组单个元素的长度
lengthn=${#array_name[n]}
```

### Shell 注释

```sh
:<<EOF
注释内容...
注释内容...
注释内容...
EOF

# EOF 也可以使用其他符号:

:<<'
注释内容...
注释内容...
注释内容...
'

:<<!
注释内容...
注释内容...
注释内容...
!
```

## Shell 传递参数

脚本内获取参数的格式为：**\$n**

**特殊参数**

| 参数 | 说明                                                                                         |
| ---- | -------------------------------------------------------------------------------------------- |
| \$0  | 当前脚本的文件名                                                                             |
| \$n  | 传递给脚本或函数的参数。n 是一个数字，表示第几个参数。例如，第一个参数是$1，第二个参数是$2。 |
| \$#  | 传递给脚本或函数的参数个数。                                                                 |
| \$\* | 传递给脚本或函数的所有参数。                                                                 |
| \$@  | 传递给脚本或函数的所有参数。被双引号(" ")包含时，与 \$\* 稍有不同，下面将会讲到。            |
| \$?  | 上个命令的退出状态，或函数的返回值。                                                         |
| \$\$ | 当前 Shell 进程 ID。对于 Shell 脚本，就是这些脚本所在的进程 ID。                             |
| \$-  | 显示 Shell 使用的当前选项，与 set 命令功能相同。                                             |
| \$!  | 后台运行的最后一个进程的 ID 号                                                               |

**$* 和 $@ 的区别**

- $* 和 $@ 都表示传递给函数或脚本的所有参数，不被双引号(" ")包含时，都以"$1" "$2" … "\$n" 的形式输出所有参数。
- 但是当它们被双引号(" ")包含时，"$*" 会将所有的参数作为一个整体，以"$1 $2 … $n"的形式输出所有参数；"$@" 会将各个参数分开，以"$1" "$2" … "$n" 的形式输出所有参数。

```sh

for var in "$*"
do
    echo $var
done

for var in "$@"
do
    echo $var
done

# ./test.sh a b c d
# output：
# a b c d
# a
# b
# c
# d
```

## Shell 数组

数组中可以存放多个值。Bash Shell **只支持一维数组**（不支持多维数组），**初始化时不需要定义数组大小**（与 PHP 类似）。

与大部分编程语言类似，**数组元素的下标由 0 开始**。

Shell **数组用括号来表示**，**元素用"空格"符号分割开**，语法格式如下：`array_name=(value1 ... valuen)`

```sh
# 定义
my_array=(A B "C" D)
# or
array_name[0]=value0
array_name[1]=value1
array_name[2]=value2

# 读取数组元素值
${array_name[index]}

# 获取数组中的所有元素
${my_array[*]}
${my_array[@]}

# 获取数组长度
${#my_array[@]}
${#my_array[*]}

# 输出所有数组索引
${!my_array}
```

## Shell 基本运算符

Shell 和其他编程语言一样，支持多种运算符，包括：

- 算数运算符
- 关系运算符
- 布尔运算符
- 字符串运算符
- 文件测试运算符

原生 bash 不支持简单的数学运算，但是可以通过其他命令来实现，例如 **awk** 和 **expr**，**expr** 最常用。

expr 是一款表达式计算工具，使用它能完成表达式的求值操作。

- **表达式和运算符之间要有空格**，例如 2+2 是不对的，必须写成 2 + 2，这与我们熟悉的大多数编程语言不一样。
- 完整的**表达式要被 `` 包含**，注意这个字符不是常用的单引号，在 Esc 键下边。
- 乘号(\*)前边必须加反斜杠(\)才能实现乘法运算；

### 算术运算符

| 运算符 | 说明                                              | 举例                          |
| ------ | ------------------------------------------------- | ----------------------------- |
| +      | 加法                                              | `expr $a + $b` 结果为 30。    |
| -      | 减法                                              | `expr $a - $b` 结果为 -10。   |
| \*     | 乘法                                              | `expr $a \* $b` 结果为 200。  |
| /      | 除法                                              | `expr $b / $a` 结果为 2。     |
| %      | 取余                                              | `expr $b % $a` 结果为 0。     |
| =      | 赋值                                              | a=\$b 将把变量 b 的值赋给 a。 |
| ==     | 相等。用于**比较两个数字**，相同则返回 true。     | [ $a == $b ] 返回 false。     |
| !=     | 不相等。用于**比较两个数字**，不相同则返回 true。 | [ $a != $b ] 返回 true。      |

```sh
val=`expr 2 + 2`
echo "两数之和为 : $val"

a=10
b=20

val=`expr $a + $b`
echo "a + b : $val"

val=`expr $a - $b`
echo "a - b : $val"

val=`expr $a \* $b`
echo "a * b : $val"

val=`expr $b / $a`
echo "b / a : $val"

val=`expr $b % $a`
echo "b % a : $val"

if [ $a == $b ]
then
   echo "a 等于 b"
fi
if [ $a != $b ]
then
   echo "a 不等于 b"
fi
```

条件表达式要放在方括号之间，并且要有空格，例如: [$a==$b] 是错误的，必须写成 [ $a == $b ]。

### 关系运算符

关系运算符**只支持数字，不支持字符串**，除非字符串的值是数字。

| 运算符 | 说明                                                  | 举例                       |
| ------ | ----------------------------------------------------- | -------------------------- |
| -eq    | 检测两个数是否相等，相等返回 true。                   | [ $a -eq $b ] 返回 false。 |
| -ne    | 检测两个数是否不相等，不相等返回 true。               | [ $a -ne $b ] 返回 true。  |
| -gt    | 检测左边的数是否大于右边的，如果是，则返回 true。     | [ $a -gt $b ] 返回 false。 |
| -lt    | 检测左边的数是否小于右边的，如果是，则返回 true。     | [ $a -lt $b ] 返回 true。  |
| -ge    | 检测左边的数是否大于等于右边的，如果是，则返回 true。 | [ $a -ge $b ] 返回 false。 |
| -le    | 检测左边的数是否小于等于右边的，如果是，则返回 true。 | [ $a -le $b ] 返回 true。  |

### 布尔运算符

| 运算符 | 说明                                                | 举例                                     |
| ------ | --------------------------------------------------- | ---------------------------------------- |
| !      | 非运算，表达式为 true 则返回 false，否则返回 true。 | [ ! false ] 返回 true。                  |
| -o     | 或运算，有一个表达式为 true 则返回 true。           | [ $a -lt 20 -o $b -gt 100 ] 返回 true。  |
| -a     | 与运算，两个表达式都为 true 才返回 true。           | [ $a -lt 20 -a $b -gt 100 ] 返回 false。 |

### 逻辑运算符

| 运算符 | 说明       | 举例                                      |
| ------ | ---------- | ----------------------------------------- |
| &&     | 逻辑的 AND | [[ $a -lt 100 && $b -gt 100 ]] 返回 false |
| \|\|   | 逻辑的 OR  | [[ \$a -lt 100                            |  | \$b -gt 100 ]] 返回 true |

### 字符串运算符

| 运算符 | 说明                                       | 举例                     |
| ------ | ------------------------------------------ | ------------------------ |
| =      | 检测两个字符串是否相等，相等返回 true。    | [ $a = $b ] 返回 false。 |
| !=     | 检测两个字符串是否相等，不相等返回 true。  | [ $a != $b ] 返回 true。 |
| -z     | 检测字符串长度是否为 0，为 0 返回 true。   | [ -z $a ] 返回 false。   |
| -n     | 检测字符串长度是否为 0，不为 0 返回 true。 | [ -n "$a" ] 返回 true。  |
| \$     | 检测字符串是否为空，不为空返回 true。      | [ $a ] 返回 true。       |

### 文件测试运算符

文件测试运算符用于检测 Unix 文件的各种属性。

| 操作符  | 说明                                                                        | 举例                      |
| ------- | --------------------------------------------------------------------------- | ------------------------- |
| -b file | 检测文件是否是块设备文件，如果是，则返回 true。                             | [ -b $file ] 返回 false。 |
| -c file | 检测文件是否是字符设备文件，如果是，则返回 true。                           | [ -c $file ] 返回 false。 |
| -d file | 检测文件是否是目录，如果是，则返回 true。                                   | [ -d $file ] 返回 false。 |
| -f file | 检测文件是否是普通文件（既不是目录，也不是设备文件），如果是，则返回 true。 | [ -f $file ] 返回 true。  |
| -g file | 检测文件是否设置了 SGID 位，如果是，则返回 true。                           | [ -g $file ] 返回 false。 |
| -k file | 检测文件是否设置了粘着位(Sticky Bit)，如果是，则返回 true。                 | [ -k $file ] 返回 false。 |
| -p file | 检测文件是否是有名管道，如果是，则返回 true。                               | [ -p $file ] 返回 false。 |
| -u file | 检测文件是否设置了 SUID 位，如果是，则返回 true。                           | [ -u $file ] 返回 false。 |
| -r file | 检测文件是否可读，如果是，则返回 true。                                     | [ -r $file ] 返回 true。  |
| -w file | 检测文件是否可写，如果是，则返回 true。                                     | [ -w $file ] 返回 true。  |
| -x file | 检测文件是否可执行，如果是，则返回 true。                                   | [ -x $file ] 返回 true。  |
| -s file | 检测文件是否为空（文件大小是否大于 0），不为空返回 true。                   | [ -s $file ] 返回 true。  |
| -e file | 检测文件（包括目录）是否存在，如果是，则返回 true。                         | [ -e $file ] 返回 true。  |

**其他检查符**：

- -S: 判断某文件是否 socket。
- -L: 检测文件是否存在并且是一个符号链接。

## Shell 流程控制

### if else

```sh
if condition1
then
    command1
elif condition2
then
    command2
else
    commandN
fi
```

### for 循环

```sh
for var in item1 item2 ... itemN
do
    command1
    command2
    ...
    commandN
done
```

### while 语句

```sh
while condition
do
    command
done
```

### until 循环

until 循环执行一系列命令直至条件为 true 时停止。

until 循环与 while 循环在处理方式上刚好相反。

一般 while 循环优于 until 循环，但在某些时候—也只是极少数情况下，until 循环更加有用。

```sh
until condition
do
    command
done
```

### case ... esac

```sh
case 值 in
模式1)
    command1
    command2
    ...
    commandN
    ;;
模式2）
    command1
    command2
    ...
    commandN
    ;;
esac
```

取值将检测匹配的每一个模式。一旦模式匹配，则执行完匹配模式相应命令后不再继续其他模式。如果无一匹配模式，使用星号 `*` 捕获该值，再执行后面的命令。

```sh
echo '输入 1 到 4 之间的数字:'
echo '你输入的数字为:'
read aNum
case $aNum in
    1)  echo '你选择了 1'
    ;;
    2)  echo '你选择了 2'
    ;;
    3)  echo '你选择了 3'
    ;;
    4)  echo '你选择了 4'
    ;;
    *)  echo '你没有输入 1 到 4 之间的数字'
    ;;
esac
```

### 跳出循环

在循环过程中，有时候需要在未达到循环结束条件时强制跳出循环，Shell 使用两个命令来实现该功能：**break**和**continue**。

## Shell 函数

```sh
[ function ] funname [()]

{

    action;

    [return int;]

}
```

1. 可以带 function fun() 定义，也可以直接 fun() 定义,不带任何参数。
2. 参数返回，可以显示加：return 返回，如果不加，将以最后一条命令运行结果，作为返回值。 return 后跟数值 n(0-255)

## Shell 输入/输出重定向

| 命令            | 说明                                               |
| --------------- | -------------------------------------------------- |
| command > file  | 将输出重定向到 file。                              |
| command < file  | 将输入重定向到 file。                              |
| command >> file | 将输出以追加的方式重定向到 file。                  |
| n > file        | 将文件描述符为 n 的文件重定向到 file。             |
| n >> file       | 将文件描述符为 n 的文件以追加的方式重定向到 file。 |
| n >& m          | 将输出文件 m 和 n 合并。                           |
| n <& m          | 将输入文件 m 和 n 合并。                           |
| << tag          | 将开始标记 tag 和结束标记 tag 之间的内容作为输入。 |

## Shell 文件包含

和其他语言一样，Shell 也可以包含外部脚本。这样可以很方便的封装一些公用的代码作为一个独立的文件。

```sh
. filename   # 注意点号(.)和文件名中间有一空格

# 或

source filename
```

且被包含的文件不需要可执行权限。

## Shell 中各种括号的作用(),(()),\[\],\[\[\]\],{}

### 小括号

#### 单小括号

1. **命令组**。括号中的命令将会新开一个子 shell 顺序执行，所以括号中的变量不能够被脚本余下的部分使用。括号中多个命令之间用分号隔开，最后一个命令可以没有分号，各命令和括号之间不必有空格。
2. **命令替换**。等同于`cmd`，shell 扫描一遍命令行，发现了`$(cmd)`结构，便将`$(cmd)`中的 cmd 执行一次，得到其标准输出，再将此输出放到原来命令。有些 shell 不支持，如 tcsh。
3. 用于**初始化数组**。如：array=(a b c d)

#### 双小括号

1. **整数扩展**。这种扩展计算是整数型的计算，不支持浮点型。((exp))结构扩展并计算一个算术表达式的值，如果表达式的结果为 0，那么返回的退出状态码为 1，或者 是"假"，而一个非零值的表达式所返回的退出状态码将为 0，或者是"true"。若是逻辑判断，表达式 exp 为真则为 1,假则为 0。
2. 只要括号中的运算符、表达式符合 C 语言运算规则，都可用在$((exp))中，甚至是三目运算符。作不同进位(如二进制、八进制、十六进制)运算时，输出结果全都自动转化成了十进制。如：echo $((16#5f)) 结果为 95 (16 进位转十进制)
3. 单纯用 (( )) 也可重定义变量值，比如 a=5; ((a++)) 可将 \$a 重定义为 6
4. 常用于算术运算比较，双括号中的变量可以不使用`$`符号前缀。括号内支持多个表达式用逗号分开。 只要括号中的表达式符合 C 语言运算规则,比如可以直接使用 for((i=0;i<5;i++)), 如果不使用双括号, 则为 for i in `seq 0 4`或者 for i in {0..4}。再如可以直接使用 if (($i<5)), 如果不使用双括号, 则为if [ $i -lt 5 ]。

### 中括号

#### 单中括号

1. **bash 的内部命令，[和 test 是等同的**。如果我们不用绝对路径指明，通常我们用的都是 bash 自带的命令。if/test 结构中的左中括号是调用 test 的命令标识，右中括号是关闭条件判断的。这个命令把它的参数作为比较表达式或者作为文件测试，并且根据比较的结果来返回一个退出状态码。if/test 结构中并不是必须右中括号，但是新版的 Bash 中要求必须这样。
2. Test 和[]中可用的比较运算符只有==和!=，两者都是用于字符串比较的，不可用于整数比较，整数比较只能使用-eq，-gt 这种形式。无论是字符串比较还是整数比较都不支持大于号小于号。如果实在想用，对于字符串比较可以使用转义形式，如果比较"ab"和"bc"：[ ab \< bc ]，结果为真，也就是返回状态为 0。[ ]中的逻辑与和逻辑或使用-a 和-o 表示。
3. **字符范围**。用作正则表达式的一部分，描述一个匹配的字符范围。作为 test 用途的中括号内不能使用正则。
4. 在一个 array 结构的上下文中，中括号用来引用数组中每个元素的编号。

#### 双中括号

1. [[是 bash 程序语言的关键字。并不是一个命令，`[[ ]]`结构比`[ ]`结构更加通用。在`[[`和`]]`之间所有的字符都不会发生文件名扩展或者单词分割，但是会发生参数扩展和命令替换。
2. 支持字符串的模式匹配，使用=~操作符时甚至支持 shell 的正则表达式。字符串比较时可以把右边的作为一个模式，而不仅仅是一个字符串，比如[[ hello == hell? ]]，结果为真。[[ ]] 中匹配字符串或通配符，不需要引号。
3. 使用[[ ... ]]条件判断结构，而不是[ ... ]，能够防止脚本中的许多逻辑错误。比如，&&、||、<和> 操作符能够正常存在于[[ ]]条件判断结构中，但是如果出现在[ ]结构中的话，会报错。比如可以直接使用 if [[ $a != 1 && $a != 2 ]], 如果不适用双括号, 则为 if [ $a -ne 1] && [ $a != 2 ]或者 if [ $a -ne 1 -a $a != 2 ]。
4. bash 把双中括号中的表达式看作一个单独的元素，并返回一个退出状态码。

### 大括号

#### 常规用法

- 大括号拓展。(通配(globbing))将对大括号中的文件名做扩展。在大括号中，不允许有空白，除非这个空白被引用或转义。第一种：对大括号中的以逗号分割的文件列表进行拓展。如 touch {a,b}.txt 结果为 a.txt b.txt。第二种：对大括号中以点点（..）分割的顺序文件列表起拓展作用，如：touch {a..d}.txt 结果为 a.txt b.txt c.txt d.txt
- 代码块，又被称为内部组，这个结构事实上创建了一个匿名函数 。与小括号中的命令不同，大括号内的命令不会新开一个子 shell 运行，即脚本余下部分仍可使用括号内变量。括号内的命令间用分号隔开，最后一个也必须有分号。{}的第一个命令和左括号之间必须要有一个空格。

#### 几种特殊的替换结构

${var:-string},${var:+string},${var:=string},${var:?string}

- `${var:-string}和${var:=string}`:若变量 var 为空，则用在命令行中用 string 来替换`${var:-string}`，否则变量 var 不为空时，则用变量 var 的值来替换`${var:-string}`；对于`${var:=string}`的替换规则和`${var:-string}`是一样的，所不同之处是`${var:=string}`若 var 为空时，用 string 替换`${var:=string}`的同时，把 string 赋给变量 var： `${var:=string}`很常用的一种用法是，判断某个变量是否赋值，没有的话则给它赋上一个默认值。
- `${var:+string}`的替换规则和上面的相反，即只有当 var 不是空的时候才替换成 string，若 var 为空时则不替换或者说是替换成变量 var 的值，即空值。(因为变量 var 此时为空，所以这两种说法是等价的)
- `${var:?string}`替换规则为：若变量 var 不为空，则用变量 var 的值来替换\${var:?string}；若变量 var 为空，则把 string 输出到标准错误中，并从脚本中退出。我们可利用此特性来检查是否设置了变量的值。
- 补充扩展：在上面这五种替换结构中 string 不一定是常值的，可用另外一个变量的值或是一种命令的输出。

#### 四种模式匹配替换结构

模式匹配记忆方法：

- \# 是去掉左边\(在键盘上\#在\$之左边\)
- % 是去掉右边(在键盘上%在\$之右边)
- #和%中的单一符号是最小匹配，两个相同符号是最大匹配。

**\${var%pattern},\${var%%pattern},\${var#pattern},\${var##pattern}**

- 第一种模式：\${variable%pattern}，这种模式时，shell 在 variable 中查找，看它是否一给的模式 pattern 结尾，如果是，就从命令行把 variable 中的内容去掉右边最短的匹配模式
- 第二种模式： \${variable%%pattern}，这种模式时，shell 在 variable 中查找，看它是否一给的模式 pattern 结尾，如果是，就从命令行把 variable 中的内容去掉右边最长的匹配模式
- 第三种模式：\${variable#pattern} 这种模式时，shell 在 variable 中查找，看它是否一给的模式 pattern 开始，如果是，就从命令行把 variable 中的内容去掉左边最短的匹配模式
- 第四种模式： \${variable##pattern} 这种模式时，shell 在 variable 中查找，看它是否一给的模式 pattern 结尾，如果是，就从命令行把 variable 中的内容去掉右边最长的匹配模式
- 这四种模式中都不会改变 variable 的值，其中，只有在 pattern 中使用了\*匹配符号时，%和%%，#和##才有区别。结构中的 pattern 支持通配符，\*表示零个或多个任意字符，?表示仅与一个任意字符匹配，[...]表示匹配中括号里面的字符，[!...]表示不匹配中括号里面的字符。

```sh
# var=testcase
# echo $var
testcase
# echo ${var%s*e}
testca
# echo $var
testcase
# echo ${var%%s*e}
te
# echo ${var#?e}
stcase
# echo ${var##?e}
stcase
# echo ${var##*e}

# echo ${var##*s}
e
# echo ${var##test}
```

#### 字符串提取和替换

**\${var:num},\${var:num1:num2},\${var/pattern/pattern},\${var//pattern/pattern}**

- 第一种模式：\${var:num}，这种模式时，shell 在 var 中提取第 num 个字符到末尾的所有字符。若 num 为正数，从左边 0 处开始；若 num 为负数，从右边开始提取字串，但必须使用在冒号后面加空格或一个数字或整个 num 加上括号，如\${var: -2}、\${var:1-3}或\${var:(-2)}。
- 第二种模式：\${var:num1:num2}，num1 是位置，num2 是长度。表示从\$var 字符串的第\$num1 个位置开始提取长度为\$num2 的子串。不能为负数。
- 第三种模式：\${var/pattern/pattern}表示将 var 字符串的第一个匹配的 pattern 替换为另一个 pattern。。
- 第四种模式：\${var//pattern/pattern}表示将 var 字符串中的所有能匹配的 pattern 替换为另一个 pattern。

```sh
[root@centos ~]# var=/home/centos
[root@centos ~]# echo $var
/home/centos
[root@centos ~]# echo ${var:5}
/centos
[root@centos ~]# echo ${var: -6}
centos
[root@centos ~]# echo ${var:(-6)}
centos
[root@centos ~]# echo ${var:1:4}
home
[root@centos ~]# echo ${var/o/h}
/hhme/centos
[root@centos ~]# echo ${var//o/h}
/hhme/cenths
```

### 符号\$后的括号

- \${a} 变量 a 的值, 在不引起歧义的情况下可以省略大括号。

- \$(cmd) 命令替换，和`cmd`效果相同，结果为 shell 命令 cmd 的输出，不过某些 Shell 版本不支持\$()形式的命令替换, 如 tcsh。

- \$((expression)) 和`exprexpression`效果相同, 计算数学表达式 exp 的数值, 其中 exp 只要符合 C 语言的运算规则即可, 甚至三目运算符和逻辑表达式都可以计算。

### 使用

**多条命令执行**

1. 单小括号，(cmd1;cmd2;cmd3) **新开一个子 shell 顺序执行命令** cmd1,cmd2,cmd3, 各命令之间用分号隔开, 最后一个命令后可以没有分号。

2. 单大括号，{ cmd1;cmd2;cmd3;} **在当前 shell 顺序执行命令** cmd1,cmd2,cmd3, 各命令之间用分号隔开, 最后一个命令后必须有分号, 第一条命令和左括号之间必须用空格隔开。

对{}和()而言, 括号中的重定向符只影响该条命令， 而括号外的重定向符影响到括号中的所有命令。

## shell 面试题

**如何调试 bash 脚本**

将 -xv 参数加到 #!/bin/bash 后

```sh
#!/bin/bash –xv
```

---

**如何进行两个整数相加?**

```sh
V1=1
V2=2
let V3=$V1+$V2
echo $V3

# 其他方法
A=5
B=6
echo $(($A+$B)) # 方法 2
echo $[$A+$B]  # 方法 3
expr $A + $B   # 方法 4
echo $A+$B | bc # 方法 5
awk 'BEGIN{print '"$A"'+'"$B"'}'  # 方法 6
```

---

**命令“export”有什么用**

使变量在子 shell 中可用

---

**在后台运行脚本**

```sh
nohup command &
```

---

**&和&&的区别**

- **&**：希望脚本在后台运行时使用它
- **&&**：当一个脚本成功完成才执行后边的（命令/脚本）的时候使用它

---

**echo \${new:-variable} 的输出是什么**

variable

---

**如何在脚本文件中重定向标准输出和标准错误流到 log.txt 文件 ?**

在脚本文件中添加**exec >log.txt 2>&1**命令。

---

**如何只用 echo 命令获取字符串变量的一部分 ?**

```sh
echo ${variable:x:y}
# x - 起始位置
# y - 长度
```

---

**如果给定字符串 variable="User:123:321:/home/dir"，如何只用 echo 命令获取 home_dir ?**

`快速记忆`：#￥%(按照键盘的字符顺序)，##和%%是最大匹配，#和%是最小匹配

```sh
echo ${variable#*:*:*:}
# or
echo ${variable##*:}
```

**如何从上面的字符串中获取 “User” ?**

```sh
echo ${variable%%:*}
# or
echo ${variable%:*:*:*}
```

**如何使用 awk 列出 UID 小于 100 的用户 ?**

```sh
awk -F: '$3<100' /etc/passwd
```

---

**写程序为用户计算主组数目并显示次数和组名**

```sh
cat /etc/passwd|cut -d: -f4|sort|uniq -c|while read c g
do
{ echo $c; grep :$g: /etc/group|cut -d: -f1;}|xargs -n 2
done
```

---

**如何在 bash shell 中更改标准的域分隔符为 ":" ?**

IFS=":"

---

**如何获取变量长度 ?**

\${#variable}

---

**如何打印变量的最后 5 个字符 ?**

echo \${variable: -5}

**${variable:-10} 和 ${variable: -10} 有什么区别?**

- \${variable:-10} - 如果之前没有给 variable 赋值则输出 10；如果有赋值则输出该变量
- \${variable: -10} - 输出 variable 的最后 10 个字符

---

**如何只用 echo 命令替换字符串的一部分 ?**

```sh
# 替换全部的匹配
echo ${variable//pattern/replacement}
# 替换第一次的匹配
echo ${variable/pattern/replacement}
```

---

**哪个命令将命令替换为大写 ?**

```sh
tr '[:lower:]' '[:upper:]'
```

---

**如何计算本地用户数目 ?**

```sh
wc -l /etc/passwd|cut -d" " -f1 或者 cat /etc/passwd|wc -l
```

---

**不用 wc 命令如何计算字符串中的单词数目 ?**

```sh
set ${string}
echo $#
```

---

**"export \$variable" 或 "export variable" 哪个正确 ?**

export variable

---

**如何列出第二个字母是 a 或 b 的文件 ?**

```sh
ls -d ?[ab]*
```

---

**如何将整数 a 加到 b 并赋值给 c ?**

```sh
c=(($a+$b))

# or
c=`expr $a + $b`

# or
c=`echo "$a+$b"|bc`
```

---

**如何去除字符串中的所有空格 ?**

```sh
echo string | tr -d " "
```

---

**写出输出数字 0 到 100 中 3 的倍数(0 3 6 9 …)的命令 ?**

```sh
for i in (0..100..3);do echo $i;done

# or

for((i=0;i<=100;i+=3));do echo $i;done
```

---

**[ $a == $b ] 和 [ $a -eq $b ] 有什么区别**

- [ $a == $b ] 用于字符串比较
- [ $a -eq $b ] 用于数字比较

---

**[[ $string == abc* ]] 和 [[ $string == "abc*" ]] 有什么区别**

- [[ $string == abc* ]] - 检查字符串是否以字母 abc 开头
- [[ $string == "abc*" ]] - 检查字符串是否完全等于 abc

---

**如何列出以 ab 或 xy 开头的用户名 ?**

egrep "^ab|^xy" /etc/passwd|cut -d: -f1

---

**$* 和 $@ 有什么区别**

- \$\* - 以一个字符串形式输出所有传递到脚本的参数
- $@ - 以 $IFS 为分隔符列出所有传递到脚本中的参数

---

**shell 脚本如何获取输入的值 ?**

```sh
# 通过参数
./script param1 param2

# 通过read命令
read -p "Destination backup Server : " desthost
```

---

**在脚本中如何使用 "expect" ?**

```sh
/usr/bin/expect << EOD
spawn rsync -ar ${line} ${desthost}:${destpath}
expect "*?assword:*"
send "${password}\r"
expect eof
EOD
```

---

## 参考

> - [菜鸟教程-shell 教程](https://www.runoob.com/linux/linux-shell.html)
> - [shell 中各种括号的作用()、(())、\[\]、\[\[\]\]、{}](https://blog.csdn.net/taiyang1987912/article/details/39551385)
> - [分享 70 个经典的 Shell 脚本面试题与答案](https://www.jb51.net/article/135168.htm)
