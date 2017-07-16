FROM schickling/latex
MAINTAINER Romano Kleinwächter "romano.kleinwaechter@gmail.com"

RUN mkdir -p /usr/local/share/render-daemon/conf

ADD daemon /usr/local/share/render-daemon
ADD conf/app.conf /usr/local/share/render-daemon/conf

RUN chmod +x /usr/local/share/render-daemon/daemon

ENTRYPOINT ["/usr/local/share/render-daemon/daemon"]

