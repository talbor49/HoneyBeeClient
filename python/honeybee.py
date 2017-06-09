import socket


class Request:
    GET_REQUEST = 1
    SET_REQUEST = 2
    DELETE_REQUEST = 3
    AUTH_REQUEST = 4
    CREATE_REQUEST = 5
    USE_REQUEST = 6
    QUIT_REQUEST = 7

    REQUEST_STATUS_KEY = 1
    REQUEST_STATUS_BUCKET = 2

    REQUEST_STATUS = 1

    RESPONSE_SIZE = 1024

    LATIN_ENCODING = 'latin-1'

    def __init__(self, request_type, status, content):
        self.type = request_type
        self.status = status
        self.content = content if isinstance(content, list) else [content, ]

    def buffer(self):
        buff = bytearray([self.type, self.status])
        buff.extend(map(ord, '\0'.join(self.content)))
        return buff

    def sbuferr(self):
        return self.buffer().decode(Request.LATIN_ENCODING)


class HoneyBee:
    def __init__(self, host, port, username, password):
        self.host = host
        self.port = port
        self.bucket = ''
        self.username = ''

        conn = socket.socket(family=socket.AF_INET)
        conn.connect((host, port))

        self.conn = conn

        self.auth(username, password)

    def __request__(self, type, status, content):
        buff = Request(type, status, content).buffer()
        self.conn.send(buff)
        print("Sending:" + buff.decode(Request.LATIN_ENCODING))
        return self.conn.recv(Request.RESPONSE_SIZE)

    def auth(self, username, password):
        return self.__request__(Request.AUTH_REQUEST, Request.REQUEST_STATUS, [username, password])

    def set(self, key, value):
        self.__request__(Request.SET_REQUEST, Request.REQUEST_STATUS, [key, value])

    def get(self, key):
        self.__request__(Request.GET_REQUEST, Request.REQUEST_STATUS, [key])

    def use(self, bucket):
        self.__request__(Request.USE_REQUEST, Request.REQUEST_STATUS, [bucket])

    def delete_key(self, key):
        self.__request__(Request.DELETE_REQUEST, Request.REQUEST_STATUS_KEY, [key])

    def delete_bucket(self, bucket):
        self.__request__(Request.DELETE_REQUEST, Request.REQUEST_STATUS_BUCKET, [bucket])

    def create(self, bucket):
        self.__request__(Request.CREATE_REQUEST, Request.REQUEST_STATUS, [bucket])

    def quit(self):
        self.__request__(Request.QUIT_REQUEST, Request.REQUEST_STATUS, [bucket])


