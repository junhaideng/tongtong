import React from "react";
import { PropTypes } from "prop-types";
import { Tabs } from "antd";
const { TabPane } = Tabs;

// tab 映射
const TabMap = [
  {
    from: "jwc",
    name: "教务处",
  },
  {
    from: "net",
    name: "网络信息中心",
  },
  {
    from: "www",
    name: "交大主页",
  },
  {
    from: "lib",
    name: "图书馆",
  },
  {
    from: "graduate",
    name: "研招网",
  },
];

// 顶部标签
class Tab extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      mode: "top",
    };
  }
  render() {
    const { mode } = this.state;
    return (
      <>
        {/* 顶部的标签页 */}
        <Tabs
          // 当前选中的标签
          defaultActiveKey={TabMap ? TabMap[0].from : ""}
          tabPosition={mode}
          onTabClick={this.props.setTab}
        >
          {/* 渲染所有的标签 */}
          {TabMap.map((item, index) => (
            <TabPane tab={item.name} key={item.from}></TabPane>
          ))}
        </Tabs>
      </>
    );
  }
}

Tab.propTypes = {
  // 设置tab，点击的时候会发生改变
  setTab: PropTypes.func.isRequired,
};

export default Tab;
