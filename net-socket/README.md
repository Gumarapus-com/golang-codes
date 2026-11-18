# About

Tutorial link: [Link](https://gumarapus.com/blog/socket-programming-with-golang)


## Usage

### First Mode

One way comunication

```sh
$ go run server0.go

<< connected to 127.0.0.1:60594
<< Hello

>> Thank you
```

```sh
$ go run client0.go

<< connected to 127.0.0.1:12345
>> Hello

<< Thank you
```


### Second Mode

Server may accept multiple connection

```sh
$ go run server1.go 
[127.0.0.1:40366] << connected
[127.0.0.1:40366] << Hello
[127.0.0.1:40366] >> Thank you

[127.0.0.1:40376] << connected
[127.0.0.1:40376] << Hello
[127.0.0.1:40376] >> Thank you

[127.0.0.1:40378] << connected
[127.0.0.1:40378] << Hello
[127.0.0.1:40378] >> Thank you
```

Client (Run multiple times)

```sh
$ go run client0.go 
<< connected to 127.0.0.1:12345
>> Hello

<< Thank you

$ go run client0.go 
...
```

### Third Mode: Chat App


Server

```sh
$ go run server2.go 
Listening on localhost:12345
Connection added 127.0.0.1:48416
Connections: [ 127.0.0.1:48416  ]
127.0.0.1:48416 << Tes from client 1

Connection added 127.0.0.1:42108
Connections: [ 127.0.0.1:48416 127.0.0.1:42108  ]
127.0.0.1:42108 << Tes from client 2
```

Client 1
```sh
$ go run client2.go 
connected
>> Tes from client 1
>> 127.0.0.1:42108 >> Tes from client 2
```

Client 2
```sh
$ go run client2.go 
connected
>> Tes from client 2
>>
```
