<!-- EventDetail.vue -->

<template>
  <div id="eventDetail" class="main-stage">
    <h2>イベント詳細</h2>
    <div class="to-edit">
        <span>
            自分で作成したイベントです
            <router-link to="/events/1/edit">
                <button>編集する</button>
            </router-link>
        </span>
    </div>
    <div class="event">
        <h3>{{ event.title }}</h3>
        <hr>
        <p>カテゴリー: {{ event.category }}</p>
        <p>雰囲気: {{ event.atmosphere }}</p>
        <p>タグ: <span v-for="tag in event.tags" :key="tag">{{ tag }}, </span></p>
        <p>説明: {{ event.description }}</p>
        <hr>
        <div class="content">
            <p>
                {{ event.content }}
            </p>
        </div>
    </div>
    <div class="bottom-content">
        <div class="comments">
            <div v-for="comment in comments" :key="comment" class="comment">
                <p>{{ comment.name }}: <span class="date">{{ comment.date }}</span></p>
                <p>{{ comment.content }}</p>
            </div>
        </div>
        <div class="request-form">
            <form action="">
                <select name="" id="">
                    <option value="0">コメントする</option>
                    <option value="1">チャットで聞く</option>
                </select>
                <textarea name="" id="" cols="30" rows="10"></textarea>
                <input type="submit">
            </form>
        </div>
    </div>
  </div>
</template>

<script>
export default {
  //props: ['event'],
  // propsで渡されたイベントデータを表示
  data() {
    return {
        event: { 
            id: 1, 
            title: 'イベント 1' ,
            category: '飲み会',
            atmosphere: 'しっぽり',
            tags: ['飲み会', 'しっぽり', '友達づくり'],
            description: 'こんなイベントです！',
            content: 'こんなイベントです。 こんなイベントです。こんなイベントです。'
        },
        comments: [
            {
                id: 1,
                name: 'ユーザー 1',
                content: 'どういうイベントですか？',
                date: '2023/10/15',
            },
            {
                id: 2,
                name: 'ユーザー 2',
                content: '一人での参加ＯＫですか？',
                date: '2023/10/15',
            },
            {
                id: 2,
                name: 'ユーザー 2',
                content: '一人での参加ＯＫですか？',
                date: '2023/10/15',
            },
            {
                id: 2,
                name: 'ユーザー 2',
                content: '一人での参加ＯＫですか？',
                date: '2023/10/15',
            },
            {
                id: 2,
                name: 'ユーザー 2',
                content: '一人での参加ＯＫですか？',
                date: '2023/10/15',
            },
            {
                id: 2,
                name: 'ユーザー 2',
                content: '一人での参加ＯＫですか？',
                date: '2023/10/15',
            },
            {
                id: 2,
                name: 'ユーザー 2',
                content: '一人での参加ＯＫですか？',
                date: '2023/10/15',
            },
        ]
    }
  }
};
</script>

<style scoped>
    .to-edit {
        text-align: right;
        margin: 10px;
        margin-left: auto;
    }

    .to-edit span {
        color: #666;
    }

    .bottom-content {
        display: flex;
    }

    .comments {
        width: 100%;
        height: 330px;
        border: solid 1px #666;
        border-radius: 5px;
        overflow-y: auto;
        padding: 10px 20px;
        margin-top: 20px; 
        margin-right: 10px;
    }

    .comment {
        margin: 10px 0;
    }

    .comment p {
        margin-top: 3px;
    }

    .date {        
        font-size: 13px;
        color: #666;
    }

    .request-form {
        width: 300px;
        border: solid 1px #666;
        border-radius: 5px;
        padding: 20px;
        margin-top: 20px;
        margin-left: auto;
    }

    select {
        border: solid 1px #666;
        border-radius: 5px;
        padding: 7px 10px;
        margin: 10px 0;
    }

    textarea {
        border: solid 1px #666;
        border-radius: 5px;
    }

    input[type="submit"] {
        background-color: #666;
        color: #fff;
        font-size: 13px;
        border-radius: 5px;
        padding: 7px 10px;
        margin-top: 10px;
        margin-left: auto;
        display: block;
    }
</style>
