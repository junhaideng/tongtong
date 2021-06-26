import React from "react";
import "../style/Notices.css";
import { PropTypes } from "prop-types";
import { Badge } from "antd";

// 消息通知信息
class Notices extends React.Component {
  render() {
    const { data } = this.props;
    return (
      <>
        {/* 渲染出所有的消息 */}
        {data.map((item, index) => (
          <div className={"notice-message"} key={index}>
            <div className={"notice-message-title"}>
              <a
                className={"notice-message-link"}
                href={item.link}
                target="_blank"
                rel={"noreferrer"}
              >
                {/* 如果是今天的消息，需要加一个badge 进行提示 */}
                {item.title}{" "}
                {item.is_today ? <Badge status={"processing"} /> : <></>}
              </a>
            </div>
            {/* 消息发布的时间 */}
            <div className="notice-message-pubDate">{item.pub_date}</div>
          </div>
        ))}
      </>
    );
  }
}

// 传递过来的数据
Notices.PropType = {
  data: PropTypes.array,
};

export default Notices;
