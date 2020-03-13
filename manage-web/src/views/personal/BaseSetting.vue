<template>
  <div class="account-settings-info-view">

    <a-row :gutter="16">
      <a-col :md="24" :lg="16">

        <a-form layout="vertical">
          <a-form-item
            label="姓名"
            :labelCol="labelCol2"
            :wrapperCol="wrapperCol2" >
            <b>{{ user.Name }}</b>
          </a-form-item>
          <a-form-item
            label="角色"
            :labelCol="labelCol2"
            :wrapperCol="wrapperCol2" >
            <b color="green">{{ userType }}</b>
          </a-form-item>
          <a-form-item
            label="权限"
            :labelCol="labelCol2"
            :wrapperCol="wrapperCol2" >
            <a-tag color="purple" v-for="item in roles" :key="item">{{ item }}</a-tag>
          </a-form-item>
        </a-form>
      </a-col>
    </a-row>
  </div>
</template>

<script>

export default {
  data () {
    return {
      labelCol2: {
        xs: { span: 24 },
        sm: { span: 5 }
      },
      wrapperCol2: {
        xs: { span: 24 },
        sm: { span: 15 }
      },
      user: {},
      roles: [],
      userType: null
    }
  },
  computed: {
    userInfo () {
      return this.$store.getters.userInfo
    }
  },
  created () {
    this.user = this.userInfo
    this.roles = this.user.SysRoles.map(item => item.name)
    switch (this.user.UserType) {
      case 'sec':
        this.userType = '教务'
        break
      case 'teacher':
        this.userType = '教师'
        break
      case 'student':
        this.userType = '学生'
        break
    }
  }
}
</script>
