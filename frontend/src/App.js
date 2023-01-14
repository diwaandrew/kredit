
import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Sidebars from './components/menu/Sidebars';
import Checklist from './components/checklist/Checklist';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Sidebars />} exact />
        <Route path="/approval" element={<Checklist />} exact />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
