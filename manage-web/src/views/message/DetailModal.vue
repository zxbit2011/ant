<template>
  <a-modal
    title="信息详情"
    :width="1000"
    :visible="visible"
    :confirmLoading="confirmLoading"
    @cancel="handleCancel"
  >
    <a-card :bordered="false" v-show="currentStep === 0">
      <detail-list :col="1">
        <detail-list-item term="信息标题">信息标题</detail-list-item>
        <detail-list-item term="发送范围"><a-tag color="green">全部</a-tag></detail-list-item>
        <detail-list-item term="指定项目"><a-tag color="blue">项目一</a-tag>  <a-tag color="blue">项目二</a-tag></detail-list-item>
        <detail-list-item term="发送状态">
          <a-tag color="purple">已发送</a-tag>
        </detail-list-item>
        <detail-list-item term="信息内容">这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦这是描述哦</detail-list-item>
        <detail-list-item term="创建时间">2019-02-06 12:22:34</detail-list-item>
        <detail-list-item term="最近修改">2019-02-06 12:22:34</detail-list-item>
        <detail-list-item term="信息附件">
          <div class="ant-pro-pages-list-projects-cardList">
            <a-list :loading="fileListLoading" :data-source="fileListData" :grid="{ gutter: 24, xl: 4, lg: 3, md: 3, sm: 2, xs: 1 }">
              <a-list-item slot="renderItem" slot-scope="item">
                <a-card class="ant-pro-pages-list-projects-card" hoverable>
                  <img slot="cover" :src="item.cover" :alt="item.title" />
                  <a-card-meta :title="item.title">
                  </a-card-meta>
                  <div class="cardItemContent">
                    <span>{{ item.updatedAt | fromNow }}</span>
                  </div>
                </a-card>
              </a-list-item>
            </a-list>
          </div>
        </detail-list-item>
      </detail-list>
    </a-card>
    <template slot="footer">
      <a-button key="back" @click="backward" v-if="currentStep > 0" :style="{ float: 'left' }" >上一步</a-button>
      <a-button key="cancel" @click="handleCancel">取消</a-button>
    </template>
  </a-modal>
</template>

<script>
import { PageView } from '@/layouts'
import { STable } from '@/components'
import DetailList from '@/components/tools/DetailList'
import moment from 'moment'

const DetailListItem = DetailList.Item

export default {
  name: 'QuotaModal',
  components: {
    PageView,
    DetailList,
    DetailListItem,
    STable
  },
  data () {
    return {
      visible: false,
      currentStep: 0,
      confirmLoading: false,
      fileListLoading: true,
      fileListData: []
    }
  },
  filters: {
    statusFilter (status) {
      const statusMap = {
        'processing': '进行中',
        'success': '完成',
        'failed': '失败'
      }
      return statusMap[status]
    },
    fromNow (date) {
      return moment(date).fromNow()
    }
  },
  computed: {
    title () {
      return this.$route.meta.title
    }
  },
  mounted () {
    this.getFileList()
  },
  methods: {
    detail (record) {
      console.log(record)
      this.visible = true
    },
    handleCancel () {
    // clear form & currentStep
      this.visible = false
      this.currentStep = 0
    },
    getFileList () {

    }
  }
}
</script>

<style lang="less" scoped>
  .title {
    color: rgba(0,0,0,.85);
    font-size: 16px;
    font-weight: 500;
    margin-bottom: 16px;
  }
</style>
