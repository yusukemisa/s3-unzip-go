## s3-unzip-go
こちらの写経

https://github.com/toshi0607/s3-unzipper-go

## Features
Write your project features

## How to use

```
$ go get github.com/yusukemisa/s3-unzip-go

$ aws s3 mb s3://${STACK_BUCKET}
```

## ハマりどころ
### cloudformationがわかりにくい（憤怒）
sam deployを実行したところ失敗
```
sam deploy --template-file sam.yml --stack-name stack-unzipper-lambda --capabilities CAPABILITY_IAM --parameter-overrides ZippedArtifactBucket= UnzippedArtifactBucket=
Waiting for changeset to be created..
Waiting for stack create/update to complete

Failed to create/update the stack. Run the following command
to fetch the list of events leading up to the failure
aws cloudformation describe-stack-events --stack-name stack-unzipper-lambda
make: *** [deploy] Error 255

```
原因としては環境変数にバケット名を定義してなかったためなのだが、
設定して再実行すると下記のようなエラーで失敗。
```
An error occurred (ValidationError) when calling the CreateChangeSet operation: Stack:arn:aws:cloudformation:ap-northeast-1:569131516825:stack/stack-unzipper-lambda/b32c4520-faf3-11e8-bb16-50d5ca9ff48e is in ROLLBACK_COMPLETE state and can not be updated.
make: *** [deploy] Error 255
```
StackがROLLBACK_COMPLETEだからダメと言われてる様子だがどうすりゃいいんだ

→結論としてはスタックを消せば解決。

どういうことかというとスタックの作成自体は初めのdeployでされており、そのスタックに定義されたリソースの作成に失敗したため
リソースの作成はロールバックされた、という状態になっていた。デプロイ操作は冪等じゃなかったということ。

### コンソールの関数のトリガーにsamのテンプレートで設定したイベントは表示されない。
テンプレートの設定ではs3オブジェクトのputをトリガーにキックされる定義を記載しているが
コンソール上の関数のトリガー設定には表示されない。

実際に動かしてキックされることを確認する以外でイベントの関連付けを確認する方法はあるか？

→





TODO: Write usage of your project
