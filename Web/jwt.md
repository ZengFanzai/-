# JWT

## 介绍

JSON Web Token(JWT)是一种互联网标准，基于 JSON 创建并用于某些声明的访问令牌。令牌使用私钥或公钥/私钥签名。

## 结构

- **header**

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

标识用于生成签名的算法。`HS256`表明此令牌是使用 HMAC-SHA256 签名的。使用典型的加密算法是 SHA-256(HS256)的 HMAC 算法和 SHA-256(RS256)的 RSA 算法。

- **Payload**

```json
{
  "loggedInAs": "admin",
  "iat": 1422779638
}
```

包含一组声明，JWT 规范定义了七个注册声明名称，它们是令牌中通常包含的标准字段。同时还可以自定义声明，具体取决于 token 的目的。如上，一个标准的时间声明`iat`和一个自定义声明`loggedInAs`。

- **Signature**

```txt
HMAC-SHA256(
 base64urlEncoding(header) + '.' +
 base64urlEncoding(payload),
 secret
)
```

安全验证 token。通过使用`Base64url` 对`header`和`payload`进行编码，并将他们用`.`分隔符连接起来计算签名。然后，通过 header 中指定的加密算法来计算签名。

- 这三个部分使用 Base64url Encoding 分别编码，并用`.`分隔符连接以产生 JWT：

```c
const token = base64urlEncoding(header) + '.' + base64urlEncoding(payload) + '.' + base64urlEncoding(signature)

// output:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dnZWRJbkFzIjoiYWRtaW4iLCJpYXQiOjE0MjI3Nzk2Mzh9.gzSraSYS8EXBxLN_oWnFSRgCzcmJmMjLiuyu5CSpyHI
```

## 使用

在身份验证中，当用户使用其凭据成功登录时，将返回 JWT 令牌，并且必须存储在本地(通常存储在`local storage`或`session storage`， 也可以存储在 cookies 中)。而不是在服务器中创建会话并返回 cookie 的传统方法。

每当用户想要访问受保护的路由或资源时，用户请求需要附带 JWT，典型的，在 header 的`Authorization`中使用`Bearer`模式：标头的内容示例：`Authorization: Bearer eyJhbGci...<snip>...yu5CSpyHI`

这是一种**无状态**的身份验证机制，用户状态不会保存在服务器内存中。

## 标准字段

| code | name            | description                                                                                                                                                                 |
| ---- | --------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| iss  | Issuer          | 标识发布 JWT 的主体。                                                                                                                                                       |
| sub  | Subject         | 标识 JWT 的主题。                                                                                                                                                           |
| aud  | Audience        | 标识 JWT 的目标接收人。每个打算处理 JWT 的负责人都必须在受众声明中使用一个值来标识自己。如果存在该声明，则处理该声明的委托人未使用 aud 声明中的值标识自身，则必须拒绝 JWT。 |
| exp  | Expiration Time | 标识到期时间，此后不得接受 JWT 进行处理。该值必须是 NumericDate：可以是整数或十进制，代表 1970-01-01 00：00：00Z 之后的秒数。                                               |
| nbf  | Not Before      | 标识开始接受 JWT 进行处理的时间。该值必须是 NumericDate。                                                                                                                   |
| iat  | Issued at       | 标识发布 JWT 的时间。该值必须是 NumericDate。                                                                                                                               |
| jti  | JWT ID          | 令牌的区分大小写的唯一标识符，即使在不同的发行者之间也是如此。                                                                                                              |

身份验证 header 中可以使用以下字段：

| code | name                                  | description                                                                  |
| ---- | ------------------------------------- | ---------------------------------------------------------------------------- |
| typ  | Token type                            | 如果存在，建议将其设置为`JWT`。                                              |
| cty  | Content type                          | 如果使用嵌套签名或加密，建议将其设置为`JWT`。否则，请忽略此字段。            |
| alg  | Message authentication code algorithm | 发行者可以自由设置算法以验证令牌上的签名。但是，某些受支持的算法是不安全的。 |

## 参考

- [JSON Web Token](https://en.wikipedia.org/wiki/JSON_Web_Token)
