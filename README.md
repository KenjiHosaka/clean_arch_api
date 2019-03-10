# Sample API Document

## 環境構築

### 初回
```bash
$ cd clean_arch_api
$ docker-compose build
$ docker-compose run --rm backend dep ensure
$ docker-compose up
```

### 2回目以降
```bash
$ cd clean_arch_api
$ docker-compose up
```

`http server started on [::]:8000` と表示されればOK

## API一覧
APIサーバーが立ち上がったら、
`http://localhost:8000/swagger/index.html` にアクセスすると見れます。