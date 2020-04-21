FROM golang

WORKDIR /home/uploader/

COPY config.json /home/uploader/
COPY main.go /home/uploader/
COPY resources/ /home/uploader/resources/
COPY route_main.go /home/uploader/
COPY template/ /home/uploader/template/
COPY utils.go /home/uploader/
COPY testgo /home/uploader/


CMD ["cd", "/home/uploader/"]

CMD ["ls", "/home/uploader"]
#バイナリ実行
CMD ["/home/uploader/testgo"]


EXPOSE 11180
