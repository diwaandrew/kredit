import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Home from './page/home/Home';
import Login from './page/login/Login';
import Sidebars from './components/menu/Sidebars';
import Checklist from './components/customertab/Checklist';
import Drawdown from './components/customertab/Drawdown'

function App() {
  if (localStorage.getItem("name")==="john"){
    return (
      <BrowserRouter>
        <div className='d-flex'>
          <div className=''>
              <Sidebars></Sidebars>
          </div>
          <div className='page'>
            <Routes>
              <Route path="/" element={<Home />} exact />
              <Route path="/approval" element={<Checklist />} exact />
              <Route path="/drawdown" element={<Drawdown />} exact />
            </Routes>
          </div>
        </div>
      </BrowserRouter>
    );
  }else{
    return(
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Login />} exact />
        </Routes>
      </BrowserRouter>
    );
  }
}

export default App;
