require 'socket'

server = TCPServer.open(8000)

loop {
    client = server.accept
    client.puts(Time.now.ctime)
    client.puts "server is running and up"
    client.close
}