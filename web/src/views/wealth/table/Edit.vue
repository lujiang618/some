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
          v-decorator="['occur_date']"
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
        <a-select v-decorator="['category', {rules: [{ required: true, message: '请选择类别' }], initialValue: '1'}]">
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
        <a-select v-decorator="['content', {rules: [{ required: true, message: '请选择状态' }], initialValue: '1'}]">
          <a-select-option v-for="item in categories" :key="item.id" >
            {{ item.name }}
          </a-select-option>
        </a-select>
      </a-form-item>

      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="金额"
        hasFeedback
        help="请填写一段描述"
      >
        <a-textarea :rows="5" placeholder="..." v-decorator="['金额', {rules: [{ required: true }]}]" />
      </a-form-item>

      <a-form-item
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        label="备注"
        hasFeedback
        validateStatus="error"
      >
        <a-date-picker
          style="width: 100%"
          showTime
          format="YYYY-MM-DD HH:mm:ss"
          placeholder="Select Time"
          v-decorator="['description']"
        />
      </a-form-item>

      <a-form-item
        v-bind="buttonCol"
      >
        <a-row>
          <a-col span="6">
            <a-button type="primary" html-type="submit">提交</a-button>
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
  // beforeCreate () {
  //   this.form = this.$form.createForm(this)
  // },
  created () {
    this.getCategories()
  },
  mounted () {
    this.$nextTick(() => {
      this.loadEditInfo(this.record)
    })
  },
  methods: {
    moment,
    ...mapActions(['CostCategories']),
    getCurrentData () {
      return new Date().toLocaleDateString()
    },
    handleGoBack () {
      this.$emit('onGoBack')
    },
    handleSubmit () {
      const { form: { validateFields } } = this
      validateFields((err, values) => {
        if (!err) {
          // eslint-disable-next-line no-console
          console.log('Received values of form: ', values)
        }
      })
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
      CostCategories().then(res => {
        if (res.result.data.length > 0) {
          this.categories = res.result.data
        }
      }).catch((err) => {
        console.log('category list', err)
      })
    }
  }
}
</script>
