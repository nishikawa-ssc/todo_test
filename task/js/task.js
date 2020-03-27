new Vue({
    el: '#taskapp',
    data: {
        // Task情報
        tasks: [],
        // 内容
        text: '',
        // 状態
        stat: -1,
        // 現在表示中の状態
        current: -1,
        // 状態一覧
        options: [
            { value: -1, label: '全て' },
            { value: 1, label: '未着手' },
            { value: 2, label: '対応中' },
            { value: 3, label: '完了'   }
        ],
        // true：入力済・false：未入力
        isEntered: false,
        dispModal: false
    },
    computed: {
        // タスクの状態一覧を表示する
        labels() {
            return this.options.reduce(function (a, b) {
                return Object.assign(a, { [b.value]: b.label })
            }, {})
        },
        // 表示対象のタスクを返却する
        computedTasks() {
          return this.tasks.filter((el) => {
            return this.current == -1 ? true : this.current === el.stat
          }, this)
        },
        // 入力チェック
        validate() {
            var val = 0 < this.text.length
            this.isEntered = val
            return val
        }
    },
    created: function() {
        this.fetchAll()
    },
    methods: {
        // 全てのタスクを取得する
        fetchAll() {
            axios.get('/task/fetchAll').then(res => {
                if (res.status != 200) {
                    throw new Error('レスポンスエラー')
                } else {
                    var datas = res.data
                    // サーバから取得したタスクをdataに設定する
                    this.tasks = datas
                }
            })
        },
        // １つのタスクを取得する
        fetchOne(task) {
            axios.get('/task/fetchOne', {
                params: {
                    id: task.id
                }
            }).then(res => {
                if (res.status != 200) {
                    throw new Error('レスポンスエラー')
                } else {
                    var dat = res.data
                    // 選択されたタスクのインデックスを取得する
                    var index = this.tasks.indexOf(task)
                    // spliceを使うとdataプロパティの配列の要素をリアクティブに変更できる
                    this.tasks.splice(index, 1, dat[0])
                }
            })
        },
        // タスクを登録する
        add() {
            // サーバへ送信するパラメータ
            const params = new URLSearchParams();
            params.append('text', this.text)
            params.append('stat', this.stat)

            axios.post('/task/add', params).then(res => {
                if (res.status != 200) {
                    throw new Error('レスポンスエラー')
                } else {
                    // 全タスクを取得する
                    this.fetchAll()
                    // 入力値を初期化する
                    this.initInputValue()
                }
            })
        },
        // タスクの状態を変更する
        edit(task) {
            // 対象のレコード取得
            axios.get('/task/fetchOne', {
                params: {
                    id: task.id
                }
            }).then(res => {
                if (res.status != 200) {
                    throw new Error('レスポンスエラー')
                } else {
                    var dat = res.data
                    // TODO: 取得したデータで確認画面表示
                }
            })
        },
        update(task) {
            // サーバへ送信するパラメータ
            const params = new URLSearchParams();
            params.append('id', task.id)

            axios.post('/task/update', params).then(res => {
                if (res.status != 200) {
                    throw new Error('レスポンスエラー')
                } else {
                    // 全タスクを取得する
                    this.fetchAll()
                }
            })
        },
        // タスクを削除する
        del(task) {
            // TODO: 削除確認画面を表示
            // サーバへ送信するパラメータ
            const params = new URLSearchParams();
            params.append('id', task.id)

            axios.post('/task/delete', params).then(res => {
                if (res.status != 200) {
                    throw new Error('レスポンスエラー')
                } else {
                    // 全タスクを取得する
                    this.fetchAll()
                }
            })
        },
        // 入力値を初期化する
        initInputValue() {
            this.stat = -1
            this.text = ''
        }
    }
})
