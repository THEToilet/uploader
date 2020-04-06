FROM ベースイメージの取得
# FROM golang:latest 
WORKDIR 作業ディレクトリの指定(存在しない場合は自動で作ります．)
# WORKDIR /go/src/github.com/THEToilet/uploader
ADD ローカルファイルをDockerイメージ内にコピーする(圧縮ファイルが自動解凍される)
# ADD: . .
COPY 同上(圧縮ファイルが自動解凍されない)
# COPY: . .
RUN Shellコマンドの実行
# RUN apt-get install -y vim
ENV 環境変数を定義
# ENV PORT=11180
EXPOSE 開放するポート番号
# EXPOSE 11180
ENTRYPOINT 実行するコマンドを指定
# ENTRYPOINT ["./testgo"]
