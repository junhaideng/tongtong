import './App.css';
import "antd/dist/antd.css";
import Home from './views/Home';

function App() {
    return (
      // 根目录，项目中其实就一个界面
        <div className={"sjtu-message"}>
            <Home/>
        </div>
    );
}

export default App;
