while true; do { \
  echo -ne "HTTP/1.0 200 OK\r\nContent-Length: $(wc -c <index.html)\r\n\r\n"; \
  cat index.html; } | nc -l -p 8080 ; \ 
done