<template>
  <div class="account-settings-info-view">
    <a-row :gutter="16">
      <a-col :md="24" :lg="16">
        <a-form layout="vertical">
          <a-form-item
            label="原密码"
            :labelCol="labelCol2"
            :wrapperCol="wrapperCol2" >
            <a-input placeholder="原密码"/>
          </a-form-item>
          <a-form-item
            label="新密码"
            :labelCol="labelCol2"
            :wrapperCol="wrapperCol2" >
            <a-input placeholder="新密码"/>
          </a-form-item>
          <a-form-item style="text-align: center">
            <a-button type="primary">修改</a-button>
            <a-button style="margin-left: 8px">取消</a-button>
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
      // cropper
      labelCol2: {
        xs: { span: 24 },
        sm: { span: 5 }
      },
      wrapperCol2: {
        xs: { span: 24 },
        sm: { span: 15 }
      }
    }
  },
  methods: {

  }
}
</script>
