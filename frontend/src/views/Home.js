import React from "react";
import Tab from "../components/Tab";
import "../style/Home.css";
import Notices from "../components/Notices";
import axios from "axios";
import { message, Spin } from "antd";

const ErrCode = -1;

// 主页内容
class Home extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      // 默认情况下信息来源
      tab: "jwc",
      // 获取到的消息数据
      data: [],
      // 是否加载中
      spinning: true,
    };
    this.setTab = this.setTab.bind(this);
  }
  // 组件挂载之后设置默认的tab
  // 并且加载对应的内容
  componentDidMount() {
    this.setData(this.state.tab);
  }

  // 点击tab之后需要重新设置一个新的state
  setTab(tab) {
    console.log("click: ", tab);
    this.setState(
      {
        tab: tab,
      },
      () => {
        // 设置加载中
        this.setState({
          spinning: true,
        });
        // 设置对应tab的数据
        this.setData(tab);
      }
    );
  }

  setData(tab) {
    // 构造请求的url
    let url = `${process.env.REACT_APP_API_URL}/api/notice/${tab}?limit=20`;
    axios
      .get(url)
      .then((res) => {
        // 如果请求消息失败
        if (res.data.code === ErrCode) {
          message.error(res.data.msg);
          return;
        }
        // 设置元素，取消加载中提示
        this.setState({
          data: res.data.data,
          spinning: false,
        });
      })
      .catch(() => {
        // 请求出现异常
        message.error("获取信息失败,请检查网络设施");
        // 关闭加载中提示
        this.setState({
          spinning: false,
        });
      });
  }

  render() {
    const { data, spinning } = this.state;
    return (
      <>
        <div className={"layout"}>
          <div>
            <Tab setTab={this.setTab} />
          </div>
          <div className={"content"}>
            <Spin tip={"加载中"} spinning={spinning}>
              {/* 通知内容 */}
              <Notices data={data} />
            </Spin>
          </div>
        </div>
      </>
    );
  }
}

export default Home;
