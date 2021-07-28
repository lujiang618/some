<template>
  <div>
    <a-form :form="form" @submit="handleSubmit">
      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="日期"
        hasFeedback
        validateStatus="success"
      >
        <a-date-picker
          style="width: 100%"
          showTime
          format="YYYY-MM-DD"
          placeholder="Select Time"
          @onChange="dateChange"
          getFieldDecorator="['occur_date', {rules: [{ required: true, message: '请选择日期' }]}]"
          :defaultValue="moment(getCurrentData(), 'YYYY-MM-DD')"
        />
      </a-form-item>

      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="类别"
        hasFeedback
        validateStatus="success"
      >
        <a-select v-decorator="['category', {rules: [{ required: true, message: '请选择类别' }]}]">
          <a-select-option v-for="item in categories" :key="item.id" >
            {{ item.name }}
          </a-select-option>
        </a-select>
      </a-form-item>

      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="内容"
        hasFeedback
        validateStatus="warning"
      >
        <a-input placeholder="内容" v-decorator="[ 'content', {rules: [{ required: true, message: '请输入内容' }]} ]"></a-input>
      </a-form-item>

      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="金额"
        hasFeedback
        help="请填写一段描述"
      >
        <a-input placeholder="金额" v-decorator="[ 'amount', {rules: [{ required: true, message: '请输入金额' }]} ]"></a-input>
      </a-form-item>

      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="备注"
        hasFeedback
        validateStatus="error"
      >
        <a-textarea :rows="5" placeholder="..." v-decorator="['description', {rules: [{ required: true }]}]" />
      </a-form-item>

      <a-form-item
        v-bind="buttonCol"
      >
        <a-row>
          <a-col span="6">
            <a-button type="primary" @click="handleSubmit">提交</a-button>
          </a-col>
          <a-col span="10">
            <a-button @click="handleGoBack">返回</a-button>
          </a-col>
          <a-col span="8"></a-col>
        </a-row>
      </a-form-item>
    </a-form>
  </div>
</template>

<script>
import moment from 'moment'
import pick from 'lodash.pick'
import { RouteView } from '@/layouts'
import { mapActions } from 'vuex'

export default {
  name: 'TableEdit',
  components: {
    RouteView
  },
  props: {
    record: {
      type: [Object, String],
      default: ''
    }
  },
  data () {
    return {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 5 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 12 }
      },
      buttonCol: {
        wrapperCol: {
          xs: { span: 24 },
          sm: { span: 12, offset: 5 }
        }
      },
      categories: [],
      form: this.$form.createForm(this),
      id: 0
    }
  },
  beforeCreate () {
    this.form = this.$form.createForm(this)
  },
  created () {
    this.getCategories()
  },
  mounted () {
    this.$nextTick(() => {
      this.loadEditInfo(this.record)
    })
  },
  methods: {
    ...mapActions(['CostCategories', 'CreateCost', 'UpdateCost']),
    moment,
    getCurrentData () {
      return new Date().toLocaleDateString()
    },
    dateChange (value, dateString) {
      this.setState({
          date: dateString
        })
    },
    handleGoBack () {
      this.$emit('onGoBack')
    },
    handleSubmit () {
      const { form: { validateFields }, UpdateCost, CreateCost } = this
      this.visible = true
      validateFields((err, values) => {
        if (!err) {
          values.user_id = 1
          values.occur_date = moment(values.occur_date).format('YYYY-MM-DD')
          console.log('Received values of form: ', values)
          if (this.record.id > 0) {
          UpdateCost(values)
            .then((res) => this.success(res))
            .catch(err => this.failed(err))
          } else {
            CreateCost(values)
              .then((res) => this.success(res))
              .catch(err => this.failed(err))
          }
          this.visible = false
        }
      })
      // if (this.visible) {
      //   return new Promise(resolve => {
      //     resolve(true)
      //   })
      // }
    },
    handleGetInfo () {

    },
    loadEditInfo (data) {
      const { form } = this
      // ajax
      console.log(`将加载 ${this.id} 信息到表单`)
      new Promise((resolve) => {
        setTimeout(resolve, 1500)
      }).then(() => {
        const formData = pick(data, ['no', 'callNo', 'status', 'description', 'updatedAt'])
        formData.updatedAt = moment(data.updatedAt)
        console.log('formData', formData)
        form.setFieldsValue(formData)
      })
    },
    getCategories () {
      const { CostCategories } = this
      var parameter = {}
      parameter.user_id = 1
      CostCategories(parameter).then(res => {
        if (res.result.data.length > 0) {
          this.categories = res.result.data
        }
      }).catch((err) => {
        console.log('category list', err)
      })
    },
    success (res) {
      console.log(res)
    },
    failed (err) {
      console.log(err)
    }
  }
}
</script>
