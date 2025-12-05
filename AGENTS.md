# AGENTS.md

このリポジトリで AI エージェント（Codex / ChatGPT など）が作業・レビューするときのガイドラインです

## プロジェクト概要
- 用途: AtCoder の回答・テンプレート管理
- 言語: Go

## ディレクトリ構成（ざっくり）
- `answers/` … 回答
- `templates/` … 再利用・よく使うコード
- `wips/` … 作業途中 (削除予定)

## コーディングスタイル
- Go: `go fmt` に準ずる

## テスト / 実行
- 特に必須の手順なし 必要に応じて実行結果を残す

## AI へのお願い
- コメント・説明は **日本語** で記載すること
- レビュー時の指摘にはプレフィックスを付ける: `[must]`・`[should]`・`[nits]`・`[imo]`
- 重大なバグにつながりそうな点は遠慮せず指摘する
- 変更内容とテスト結果は簡潔にまとめ、読みやすさを重視する
- changed files 以外は変更・指摘しない
- AI エージェントがコミットするときは、どのエージェントが作業したか分かるようコミットメッセージ末尾に `Co-authored-by` を付ける
  - codex: `Co-authored-by: Codex <chatgpt-codex-connector@users.noreply.github.com>`
