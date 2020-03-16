# HTTP

## 打开一个网页，整个过程会使用哪些协议(在浏览器中输入 url 地址 ->> 显示主页的过程)

![HTTP请求过程](./img/HTTP_REQ_PROCESS.png)

## HTTP 的基本性质

- HTTP 是简单的
- HTTP 是可扩展的（HTTP headers）
- HTTP 是无状态的，有会话的（可使用 Cookies 创建有状态的会话）
- HTTP 和连接
  - HTTP/1.0 默认为每一对 HTTP 请求/响应都打开一个单独的 TCP 连接
  - HTTP/1.1 引入了流水线（被证明难以实现）和持久连接的概念
  - HTTP/2 则发展得更远，通过在一个连接复用消息的方式来让这个连接始终保持为暖连接。

## HTTP 缓存

- （私有）浏览器缓存
- （共享）代理缓存

### 缓存控制(Cache-control 头)

- 禁止进行缓存：`Cache-Control: no-store`
- 强制确认缓存：`Cache-Control: no-cache`，此方式下，每次有请求发出时，缓存会将此请求发到服务器，服务器端会验证请求中所描述的缓存是否过期，若未过期（注：实际就是返回 304），则缓存才使用本地缓存副本。
- 私有缓存和公共缓存：`Cache-Control: private` 或`Cache-Control: public`。
  - "public" 指令表示该响应可以被任何中间人（中间代理、CDN 等）缓存
  - 一些通常不被中间人缓存的页面（默认是 private）（比如 带有 HTTP 验证信息（帐号密码）的页面 或 某些特定状态码的页面），将会被其缓存。

### 缓存过期机制

`Cache-Control: max-age=31536000`

### 缓存验证确认

`Cache-Control: must-revalidate`

意味着缓存在考虑使用一个陈旧的资源时，必须先验证它的状态，已过期的缓存将不被使用

### 新鲜度

一方面，由于缓存只有有限的空间用于存储资源副本，所以缓存会定期地将一些副本删除，这个过程叫做**缓存驱逐**。

另一方面，当服务器上面的资源进行了更新，那么缓存中的对应资源也应该被更新，由于 HTTP 是 C/S 模式的协议，服务器更新一个资源时，不可能直接通知客户端更新缓存，所以双方必须为该资源约定一个过期时间，在该过期时间之前，该资源（缓存副本）就是新鲜的，当过了过期时间后，该资源（缓存副本）则变为陈旧的。

驱逐算法用于将陈旧的资源（缓存副本）替换为新鲜的，注意，一个陈旧的资源（缓存副本）是不会直接被清除或忽略的，当客户端发起一个请求时，缓存检索到已有一个对应的陈旧资源（缓存副本），则缓存会先将此请求附加一个`If-None-Match`头，然后发给目标服务器，以此来检查该资源副本是否是依然还是算新鲜的，若服务器返回了 304 (Not Modified)（该响应不会有带有实体信息），则表示此资源副本是新鲜的，这样一来，可以节省一些带宽。若服务器通过 `If-None-Match` 或 `If-Modified-Since`判断后发现已过期，那么会带有该资源的实体内容返回。

![HTTPStaleness](./img/HTTPStaleness.png)

对于含有特定头信息的请求，会去计算缓存寿命。比如 Cache-control: max-age=N 的头，相应的缓存的寿命就是 N。通常情况下，对于不含这个属性的请求则会去查看是否包含 Expires 属性，通过比较 Expires 的值和头里面 Date 属性的值来判断是否缓存还有效。如果 max-age 和 expires 属性都没有，找找头里的 Last-Modified 信息。如果有，缓存的寿命就等于头里面 Date 的值减去 Last-Modified 的值除以 10（注：根据 rfc2626 其实也就是乘以 10%）。

缓存失效时间计算公式如下：

```txt
expirationTime = responseTime + freshnessLifetime - currentAge
```

上式中，responseTime 表示浏览器接收到此响应的那个时间点。

### Vary

Vary HTTP 响应头决定了对于后续的请求头，如何判断是请求一个新的资源还是使用缓存的文件。

`Vary: User-Agent`

## HTTP Cookie

Cookie 主要用于以下三个方面：

- 会话状态管理（用户登录状态、购物车、游戏分数或其它需要记录的信息）
- 个性化设置(用户自定义设置、主题等)
- 浏览器行为跟踪（如跟踪分析用户行为等）

新的浏览器 API 已经允许开发者直接将数据存储到本地，如使用 Web storage API （本地存储 localStorge 和会话存储 sessionStorge）或 IndexedDB 。

### 创建 Cookie

服务器在响应头添加`Set-Cookie`选项，浏览器会保存 cookie，并在每次请求中附带 cookie。另外，Cookie 的过期时间、域、路径、有效期、适用站点都可以根据需要来指定。

`Set-Cookie: <cookie名>=<cookie值>`

### 会话期 Cookie

浏览器关闭之后会自动删除，仅在会话期有效，不需要指定过期时间(Expires)或有效期(Max-Age)。有些浏览器提供了会话恢复功能，这种情况下即使关闭了浏览器，会话期 Cookie 也会被保留下来，就好像浏览器从来没有关闭一样。

### 持久性 Cookie

和关闭浏览器便失效的会话期 Cookie 不同，持久性 Cookie 可以指定一个特定的过期时间（Expires）或有效期（Max-Age）。

`Set-Cookie: id=a3fWa; Expires=Wed, 21 Oct 2015 07:28:00 GMT;`

### Cookie 的 Secure 和 HttpOnly 标记

标记为 Secure 的 Cookie 只应通过被 HTTPS 协议加密过的请求发送给服务端。

有 HttpOnly 标记的 Cookie,通过 JavaScript 的 Document.cookie API 无法访问

`Set-Cookie: id=a3fWa; Expires=Wed, 21 Oct 2015 07:28:00 GMT; Secure; HttpOnly`

### Cookie 的作用域

Domain 和 Path 标识定义了 Cookie 的作用域：即 Cookie 应该发送给哪些 URL。

`Set-Cookie: Domain=mozilla.org;Path= Path=/docs`

### SameSite Cookie

SameSite Cookie 允许服务器要求某个 cookie 在跨站请求时不会被发送，从而可以阻止跨站请求伪造攻击（CSRF）。

### 追踪和隐私

- 第三方 Cookie：Cookie 的域和页面的域不同，则称之为第三方 Cookie（third-party cookie.）
- 禁止追踪 Do-Not-Track
- 僵尸 Cookie 和删不掉的 Cookie，这类 Cookie 较难以删除，甚至删除之后会自动重建。它们一般是使用 Web storage API、Flash 本地共享对象或者其他技术手段来达到的

## CORS(跨域资源共享)

跨域资源共享(CORS) 是一种机制，它使用额外的 HTTP 头来告诉浏览器让运行在一个 origin (domain) 上的 Web 应用被准许访问来自不同源服务器上的指定的资源。

### 什么情况下需要 CORS

- 由 XMLHttpRequest 或 Fetch 发起的跨域 HTTP 请求。
- Web 字体 (CSS 中通过 @font-face 使用跨域字体资源), 因此，网站就可以发布 TrueType 字体资源，并只允许已授权网站进行跨站调用。
- WebGL 贴图
- 使用 drawImage 将 Images/video 画面绘制到 canvas

### 预检请求

与简单请求不同，“需预检的请求”要求必须首先使用 OPTIONS 方法发起一个预检请求到服务器，以获知服务器是否允许该实际请求。"预检请求“的使用，可以避免跨域请求对服务器的用户数据产生未预期的影响。

当请求满足下述任一条件时，即应首先发送预检请求：

下面任一 HTTP 方法：

- PUT
- DELETE
- CONNECT
- OPTIONS
- TRACE
- PATCH

人为设置了对 CORS 安全的首部字段集合之外的其他首部字段。该集合为：

- Accept
- Accept-Language
- Content-Language
- Content-Type (需要注意额外的限制)
- DPR
- Downlink
- Save-Data
- Viewport-Width
- Width

Content-Type 的值不属于下列之一：

- application/x-www-form-urlencoded
- multipart/form-data
- text/plain

请求中的 XMLHttpRequestUpload 对象注册了任意多个事件监听器。

请求中使用了 ReadableStream 对象。

![preflight_correct](./img/preflight_correct.png)

### 附带身份凭证的请求

Fetch 与 CORS 的一个有趣的特性是，可以基于 HTTP cookies 和 HTTP 认证信息发送身份凭证。一般而言，对于跨域 XMLHttpRequest 或 Fetch 请求，浏览器不会发送身份凭证信息。如果要发送凭证信息，需要设置 XMLHttpRequest 的某个特殊标志位。

```js
var invocation = new XMLHttpRequest();
var url = "http://bar.other/resources/credentialed-content/";

function callOtherDomain() {
  if (invocation) {
    invocation.open("GET", url, true);
    invocation.withCredentials = true;
    invocation.onreadystatechange = handler;
    invocation.send();
  }
}
```

第 7 行将 XMLHttpRequest 的 `withCredentials 标志设置为 true`，从而向服务器发送 Cookies。因为这是一个简单 GET 请求，所以浏览器不会对其发起“预检请求”。但是，如果`服务器端的响应中未携带 Access-Control-Allow-Credentials: true ，浏览器将不会把响应内容返回给请求的发送者`。

![cred-req](./img/cred-req.png)

**对于附带身份凭证的请求，服务器不得设置 Access-Control-Allow-Origin 的值为“*”。**

### HTTP 响应首部字段

#### Access-Control-Allow-Origin: <origin> | \*

访问控制允许来源

origin 参数的值指定了允许访问该资源的外域 URI。对于不需要携带身份凭证的请求，服务器可以指定该字段的值为通配符，表示允许来自所有域的请求。

`Access-Control-Allow-Origin: http://mozilla.com`

#### Access-Control-Expose-Headers

访问控制公开标头，让服务器把允许浏览器访问的头放入白名单

`Access-Control-Expose-Headers: X-My-Custom-Header, X-Another-Custom-Header`

这样浏览器就能够通过 getResponseHeader 访问 X-My-Custom-Header 和 X-Another-Custom-Header 响应头了。

#### Access-Control-Max-Age

访问控制最大年龄，指定了 preflight 请求的结果能够被缓存多久

`Access-Control-Max-Age: <delta-seconds>`

#### Access-Control-Allow-Credentials

访问控制允许凭证，指定了当浏览器的credentials设置为true时是否允许浏览器读取response的内容。

`Access-Control-Allow-Credentials: true`

#### Access-Control-Allow-Methods

访问控制允许方法,用于预检请求的响应。其指明了实际请求所允许使用的 HTTP 方法。

`Access-Control-Allow-Methods: <method>[, <method>]*`

#### Access-Control-Allow-Headers

访问控制允许标题,用于预检请求的响应。其指明了实际请求中允许携带的首部字段。

`Access-Control-Allow-Headers: <field-name>[, <field-name>]*`

## 状态码

![HTTP状态码](./img/HTTP_STATUS_CODE.png)

**200 OK** 从客户端发送的请求在服务端被正常处理了。

**201 Created**

**202 Accepted**

**204 No Content** 服务器接受的请求已成功处理，但返回的响应报文的主体部分不包含实体（浏览器页面不更新，仅发送信息给服务器）

**206 Partial Content** 客户端进行了范围请求，而服务器成功执行了这部分请求

**301 Moved Permanently** 永久性重定向，请求的资源已经分配了新的 URI，以后应该使用资源现在
所指的 URI

**302 Found** 临时性重定向，请求的资源临时分配了新的 URI，希望用户本次可以使用新的 URI 访问

注意：当 301,302,303 响应状态码返回时，几乎所有的浏览器会把 POST 改成 GET，并删除请求报文内的主体，之后请求会自动再次发送

**304 Not Modified** 服务器端资源未改变，可直接使用客户端未过期的缓存，不包含任何响应的主体部分

**400 Bad Request** 请求报文中存在语法错误

**401 Unauthorized** 请求需要有通过 HTTP 认证的认证信息。另外如果之前已进行一次请求，则表示用户认证失败

**403 Forbidden** 对请求资源的访问被服务器拒绝了，如未获得文件系统的访问授权，访问权限出现某些问题

**404 Not Found** 服务器上没有请求的资源

**499 client has closed connection** 服务器端处理的时间过长，客户端主动断开链接，ngix 定义的状态码

**500 Internal Server Error** 服务器在执行请求时发生了错误，Bug 或临时故障

**502 Bad Gateway** 服务器作为网关需要得到一个处理这个请求的响应，但是得到一个错误的响应。

**503 Service Unavailable** 服务器暂时处于超负荷或正在进行停机维护，现在无法处理请求

**504 Gateway Timeout**

## 各种协议与 HTTP 协议之间的关系

![各种协议与HTTP协议之间的关系](./img/HTTP_HANDLE.png)

## HTTP 长连接,短连接

在 HTTP/1.0 中默认使用短连接。也就是说，客户端和服务器每进行一次 HTTP 操作，就建立一次连接，任务结束就中断连接。当客户端浏览器访问的某个 HTML 或其他类型的 Web 页中包含有其他的 Web 资源（如 JavaScript 文件、图像文件、CSS 文件等），每遇到这样一个 Web 资源，浏览器就会重新建立一个 HTTP 会话。

而从 HTTP/1.1 起，默认使用长连接，用以保持连接特性。使用长连接的 HTTP 协议，会在响应头加入这行代码：`Connection:keep-alive`

在使用长连接的情况下，当一个网页打开完成后，客户端和服务器之间用于传输 HTTP 数据的 TCP 连接不会关闭，客户端再次访问这个服务器时，会继续使用这一条已经建立的连接。Keep-Alive 不会永久保持连接，它有一个保持时间，可以在不同的服务器软件（如 Apache）中设定这个时间。实现长连接需要客户端和服务端都支持长连接。

## HTTP 特点

1. **简单快速**：客户向服务器请求服务时，只需传送请求方法和路径。请求方法常用的有 GET、HEAD、POST。每种方法规定了客户与服务器联系的类型不同。由于 HTTP 协议简单，使得 HTTP 服务器的程序规模小，因而通信速度很快。
2. **灵活**：HTTP 允许传输任意类型的数据对象。正在传输的类型由 Content-Type 加以标记。
3. **无连接**：无连接的含义是**限制每次连接只处理一个请求**。服务器处理完客户的请求，并收到客户的应答后，即断开连接。采用这种方式可以节省传输时间。
4. **无状态**：HTTP 协议是无状态协议。无状态是指**协议对于事务处理没有记忆能力**。缺少状态意味着如果后续处理需要前面的信息，则它必须重传，这样可能导致每次连接传送的数据量增大。另一方面，在服务器不需要先前信息时它的应答就较快。
5. 支持 B/S 及 C/S 模式。

## HTTP/2 特性

## 二进制分帧

**帧**：HTTP/2 数据通信的最小单位消息：指 HTTP/2 中逻辑上的 HTTP 消息。例如请求和响应等，消息由一个或多个帧组成。

**流**：存在于连接中的一个虚拟通道。流可以承载双向消息，每个流都有一个唯一的整数 ID。

**HTTP/2 采用二进制格式传输数据，而非 HTTP 1.x 的文本格式，二进制协议解析起来更高效。** HTTP / 1 的请求和响应报文，都是由起始行，首部和实体正文（可选）组成，各部分之间以文本换行符分隔。HTTP/2 将请求和响应数据分割为更小的帧，并且它们采用二进制编码。

**HTTP/2 中，同域名下所有通信都在单个连接上完成，该连接可以承载任意数量的双向数据流。**每个数据流都以消息的形式发送，而消息又由一个或多个帧组成。多个帧之间可以乱序发送，根据帧首部的流标识可以重新组装。

## 多路复用

多路复用，代替原来的序列和阻塞机制。所有请求都是通过一个 TCP 连接并发完成。

在 HTTP/2 中，有了二进制分帧之后，HTTP /2 不再依赖 TCP 链接去实现多流并行了，在 HTTP/2 中：

- 同域名下所有通信都在单个连接上完成。
- 单个连接可以承载任意数量的双向数据流。
- 数据流以消息的形式发送，而消息又由一个或多个帧组成，多个帧之间可以乱序发送，因为根据帧首部的流标识可以重新组装。

这一特性，使性能有了极大提升：

- **同个域名只需要占用一个 TCP 连接**，消除了因多个 TCP 连接而带来的延时和内存消耗。
- 单个连接上可以并行交错的请求和响应，之间互不干扰。
- 在 HTTP/2 中，每个请求都可以带一个 31bit 的优先值，0 表示最高优先级， 数值越大优先级越低。有了这个优先值，客户端和服务器就可以在处理不同的流时采取不同的策略，以最优的方式发送流、消息和帧。

## 服务器推送

服务端可以在发送页面 HTML 时**主动推送其它资源**，而不用等到浏览器解析到相应位置，发起请求再响应。例如服务端可以主动把 JS 和 CSS 文件推送给客户端，而不需要客户端解析 HTML 时再发送这些请求。

服务端可以主动推送，客户端也有权利选择是否接收。如果服务端推送的资源已经被浏览器缓存过，浏览器可以通过发送 RST_STREAM 帧来拒收。主动推送也遵守同源策略，服务器不会随便推送第三方资源给客户端。

## 头部压缩

HTTP/2 对消息头采用 HPACK（专为 http/2 头部设计的压缩格式）进行压缩传输，能够节省消息头占用的网络的流量。而 HTTP/1.x 每次请求，都会携带大量冗余头信息，浪费了很多带宽资源。

为了减少这块的资源消耗并提升性能， **HTTP/2 对这些首部采取了压缩策略**：

- HTTP/2 在客户端和服务器端使用“首部表”来跟踪和存储之前发送的键－值对，对于相同的数据，不再通过每次请求和响应发送；
- 首部表在 HTTP/2 的连接存续期内始终存在，由客户端和服务器共同渐进地更新;
- 每个新的首部键－值对要么被追加到当前表的末尾，要么替换表中之前的值。

![HTTP2头部压缩](./img/HTTP2_compress_header.jpg)

## 常用 HTTP Header

## 请求 request

| Header          | 解释                                                                           | 示例                                              |
| --------------- | ------------------------------------------------------------------------------ | ------------------------------------------------- |
| Accept          | 指定客户端能够接收的内容类型                                                   | Accept: text/plain, text/html                     |
| Accept-Encoding | 指定浏览器可以支持的 web 服务器返回内容压缩编码类型。                          | Accept-Encoding: compress, gzip                   |
| Accept-Language | 浏览器可接受的语言                                                             | Accept-Language: en,zh                            |
| Authorization   | HTTP 授权的授权证书                                                            | Authorization: Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ== |
| Cache-Control   | 指定请求和响应遵循的缓存机制                                                   | Cache-Control: no-cache                           |
| Connection      | 表示是否需要持久连接。（HTTP 1.1 默认进行持久连接）                            | Connection: close                                 |
| Cookie          | HTTP 请求发送时，会把保存在该请求域名下的所有 cookie 值一起发送给 web 服务器。 | Cookie: \$Version=1; Skin=new;                    |
| Content-Length  | 请求的内容长度                                                                 | Content-Length: 348                               |
| Content-Type    | 请求的与实体(body)对应的 MIME 信息                                             | Content-Type: application/x-www-form-urlencoded   |
| Host            | 指定请求的服务器的域名和端口号                                                 | Host: www.baidu.com                               |
| Referer         | 先前网页的地址，当前请求网页紧随其后,即来路                                    | Referer: http://www.baidu.com/archives/71.html    |
| User-Agent      | User-Agent 的内容包含发出请求的用户信息                                        | User-Agent: Mozilla/5.0 (Linux; X11)              |

## 响应 response

| Header           | 解释                                                                                  | 示例                                                |
| ---------------- | ------------------------------------------------------------------------------------- | --------------------------------------------------- |
| Accept-Ranges    | 表明服务器是否支持指定范围请求及哪种类型的分段请求                                    | Accept-Ranges: bytes                                |
| Age              | 从原始服务器到代理缓存形成的估算时间（以秒计，非负）                                  | Age: 12                                             |
| Cache-Control    | 告诉所有的缓存机制是否可以缓存及哪种类型                                              | Cache-Control: no-cache                             |
| Allow            | 对某网络资源的有效的请求行为，不允许则返回 405                                        | Allow: GET, HEAD                                    |
| Content-Encoding | web 服务器支持的返回内容压缩编码类型。                                                | Content-Encoding: gzip                              |
| Content-Language | 响应体的语言                                                                          | Content-Language: en,zh                             |
| Content-Length   | 响应体的长度                                                                          | Content-Length: 348                                 |
| Content-Type     | 返回内容的 MIME 类型                                                                  | Content-Type: text/html; charset=utf-8              |
| Expires          | 响应过期的日期和时间                                                                  | Expires: Thu, 01 Dec 2010 16:00:00 GMT              |
| Location         | 用来重定向接收方到非请求 URL 的位置来完成请求或标识新的资源                           | Location: http://honglu.me/archives/                |
| refresh          | 应用于重定向或一个新的资源被创造，在 5 秒之后重定向（由网景提出，被大部分浏览器支持） | Refresh: 5; url=http://honglu.me/archives/          |
| Server           | web 服务器软件名称                                                                    | Server: Apache/1.3.27 (Unix) (Red-Hat/Linux)        |
| Set-Cookie       | 设置 Http Cookie                                                                      | Set-Cookie: UserID=JohnDoe; Max-Age=3600; Version=1 |

> - [常用 HTTP 状态码总结](https://www.jianshu.com/p/97f24ef4e617)
> - [一文读懂 HTTP/2 特性](https://zhuanlan.zhihu.com/p/26559480)
> - [开发中常用的 HTTP header](https://honglu.me/2015/07/13/%E5%BC%80%E5%8F%91%E4%B8%AD%E5%B8%B8%E7%94%A8%E7%9A%84HTTP-header/)
> - [HTTP](https://developer.mozilla.org/zh-CN/docs/Web/HTTP)
