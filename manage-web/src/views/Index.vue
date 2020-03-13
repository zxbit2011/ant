<template>
  <page-view :avatar="avatar" :title="false">
    <div slot="headerContent">
      <div class="title">{{ timeFix }}，{{ user.Name }}<span class="welcome-text">，{{ welcome() }}</span></div>
      <div><a-tag color="purple" v-for="item in roles" :key="item">{{ item }}</a-tag></div>
    </div>
    <div slot="extra">
      <a-row class="more-info">
        <a-col :span="6">
          <head-info title="项目数" :content="group.projectNum" :center="false" :bordered="false"/>
        </a-col>
        <a-col :span="6">
          <head-info title="指标数" :content="group.quotaNum" :center="false" :bordered="false"/>
        </a-col>
        <a-col :span="6">
          <head-info title="任务数" :content="group.taskNum" :center="false"/>
        </a-col>
        <a-col :span="6">
          <head-info title="信息通知" :content="group.notifyNum" :center="false"/>
        </a-col>
      </a-row>
    </div>

    <div>
      <a-row :gutter="24">
        <a-col :xl="16" :lg="24" :md="24" :sm="24" :xs="24">
          <a-card
            class="project-list"
            :loading="loading"
            style="margin-bottom: 24px;"
            :bordered="false"
            title="进行中的项目"
            :body-style="{ padding: 0 }">
            <a slot="extra" @click="$router.push({ name: 'projectList'})" v-if="$auth('index_project_group')">全部项目</a>
            <div>
              <a-card-grid class="project-card-grid" :key="i" v-for="(item, i) in projects">
                <a-card :bordered="false" :body-style="{ padding: 0 }">
                  <a-card-meta>
                    <div slot="title" class="card-title">
                      <a>{{ item.name }}</a>
                    </div>
                    <div slot="description" class="card-description">
                      {{ formatDate(item.start_time) }} ~ {{ formatDate(item.end_time) }}
                    </div>
                  </a-card-meta>
                  <div class="project-item">
                    <a>
                      <a-tag :color="ReviewStatusColor(item.review_status)">
                        {{ ReviewStatusText(item.review_status) }}
                      </a-tag>
                    </a>
                    <span class="datetime">{{ formatDate(item.created_at) }}</span>
                  </div>
                </a-card>
              </a-card-grid>
              <div v-if="projects.length===0" class="ant-list-empty-text" style="padding: 40px"><div class="no-data">暂无数据</div></div>
            </div>
          </a-card>
        </a-col>
        <a-col
          style="padding: 0 12px"
          :xl="8"
          :lg="24"
          :md="24"
          :sm="24"
          :xs="24">
          <a-card :loading="loading" title="最新信息" :bordered="false">
            <a slot="extra" @click="$router.push({ name: 'message'})" v-if="$auth('index_sys_notify_group')">全部信息</a>
            <a-list>
              <a-list-item :key="index" v-for="(item, index) in notifys">
                <a-list-item-meta>
                  <div slot="title">
                    <div>{{ item.title }}</div>
                  </div>
                  <div slot="description">
                    <span>
                      <a-badge :status="item.status | statusTypeFilter" :text="item.status | statusFilter"/>
                    </span>
                    <span style="float: right">
                      {{ formatDate(item.created_at) }}
                    </span>
                  </div>
                </a-list-item-meta>
              </a-list-item>
            </a-list>
          </a-card>
        </a-col>
      </a-row>
    </div>
  </page-view>
</template>

<script>
import { formatDate, timeFix } from '@/utils/util'
import { ReviewStatusColor, ReviewStatusText } from '@/utils/enum'
import { mapGetters } from 'vuex'

import { PageView } from '@/layouts'
import HeadInfo from '@/components/tools/HeadInfo'
import { GetSysNotifyList } from '@/api/notify'

const statusMap = {
  1: {
    status: 'success',
    text: '已发布'
  },
  2: {
    status: 'processing',
    text: '未发布'
  }
}
var that
export default {
  name: 'Index',
  components: {
    PageView,
    HeadInfo
  },
  data () {
    return {
      timeFix: timeFix(),
      avatar: '',
      user: {},
      roles: [],
      projects: [],
      loading: true,
      notifys: [],
      group: { },
      // data
      axis1Opts: {
        dataKey: 'item',
        line: null,
        tickLine: null,
        grid: {
          lineStyle: {
            lineDash: null
          },
          hideFirstLine: false
        }
      },
      axis2Opts: {
        dataKey: 'score',
        line: null,
        tickLine: null,
        grid: {
          type: 'polygon',
          lineStyle: {
            lineDash: null
          }
        }
      },
      scale: [{
        dataKey: 'score',
        min: 0,
        max: 80
      }]
    }
  },
  computed: {
    userInfo () {
      return this.$store.getters.userInfo
    }
  },
  filters: {
    statusFilter (type) {
      return statusMap[parseInt(type)].text
    },
    statusTypeFilter (type) {
      return statusMap[parseInt(type)].status
    }
  },
  created () {
    that = this
    this.user = this.userInfo
    this.avatar = this.userInfo.Photo
    this.roles = this.user.SysRoles.map(item => item.name)
  },
  mounted () {
    this.getNotify()
  },
  methods: {
    formatDate,
    ReviewStatusText,
    ReviewStatusColor,
    ...mapGetters(['nickname', 'welcome']),
    getNotify () {
      that.loading = true
      GetSysNotifyList({ limit: 12 }).then(res => {
        that.loading = false
        if (res.ret === 200) {
          that.notifys = res.data
        }
      })
    }
  }
}
</script>

<style lang="less" scoped>
  .project-list {

  .card-title {
    font-size: 0;

  a {
    color: rgba(0, 0, 0, 0.85);
    line-height: 24px;
    height: 24px;
    display: inline-block;
    vertical-align: top;
    font-size: 14px;

  &
  :hover {
    color: #1890ff;
  }

  }
  }
  .card-description {
    color: rgba(0, 0, 0, 0.45);
    height: 44px;
    line-height: 22px;
    overflow: hidden;
  }

  .project-item {
    display: flex;
    margin-top: 8px;
    overflow: hidden;
    font-size: 12px;
    height: 22px;
    line-height: 20px;

  a {
    color: rgba(0, 0, 0, 0.45);
    display: inline-block;
    flex: 1 1 0;

  &
  :hover {
    color: #1890ff;
  }

  }
  .datetime {
    color: rgba(0, 0, 0, 0.25);
    flex: 0 0 auto;
    float: right;
  }

  }
  .ant-card-meta-description {
    color: rgba(0, 0, 0, 0.45);
    height: 44px;
    line-height: 22px;
    overflow: hidden;
  }

  }

  .item-group {
    padding: 20px 0 8px 24px;
    font-size: 0;

  a {
    color: rgba(0, 0, 0, 0.65);
    display: inline-block;
    font-size: 14px;
    margin-bottom: 13px;
    width: 25%;
  }

  }

  .members {

  a {
    display: block;
    margin: 12px 0;
    line-height: 24px;
    height: 24px;

  .member {
    font-size: 14px;
    color: rgba(0, 0, 0, .65);
    line-height: 24px;
    max-width: 100px;
    vertical-align: top;
    margin-left: 12px;
    transition: all 0.3s;
    display: inline-block;
  }

  &
  :hover {

  span {
    color: #1890ff;
  }

  }
  }
  }

  .mobile {

  .project-list {

  .project-card-grid {
    width: 100%;
  }

  }

  .more-info {
    border: 0;
    padding-top: 16px;
    margin: 16px 0 16px;
  }

  .headerContent .title .welcome-text {
    display: none;
  }

  }

</style>
