#!/bin/sh

# スキーマのディレクトリ名
schema_dir="pkg/schema"

dir=$(PWD)

# --initでローカルの設定ファイルを初期化、updateでデータの取得と親プロジェクトで指定されているコミットにチェックアウトする。
git submodule update --init

# submoduleとして取り込んであるschemaリポジトリにsparsecheckoutを設定する
echo "set $schema_dir to sparse-checkout"

mkdir -p ./.git/modules/${schema_dir}/info
echo '/gen/server' >./.git/modules/${schema_dir}/info/sparse-checkout
cd ${schema_dir}
git config core.sparsecheckout true

# # sparsecheckoutの設定を反映
git read-tree -mu HEAD
cd ${dir}
