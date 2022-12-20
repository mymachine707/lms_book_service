1. create database book_service in your postgres!
# Migrate Db up:
# ```migrate -path ./storage/migrations -database 'postgres://user:password@127.0.0.1:5432/book_service?sslmode=disable' up```

2. cd lms_proto
    git pull
    cd ..
    ./genproto.sh