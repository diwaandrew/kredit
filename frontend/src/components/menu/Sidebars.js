import React, { useState } from 'react'
import { Link } from 'react-router-dom';
import { Menu } from './menu';
import bsim from '../../assets/favicon.ico'
// import SidebarMenu from 'react-bootstrap-sidebar-menu';
import {Image} from 'react-bootstrap';
import { IconContext } from 'react-icons';
import * as FcIcons from "react-icons/fc";
import './Sidebars.css'

const Sidebars = ({ selectData }) => {
    const [sidebar, setsidebar] = useState(true)

    const showSidebar = () => {
        setsidebar(!sidebar)
        selectData(!sidebar)
    }
    const logout = () => {
        localStorage.removeItem('nik');
        localStorage.removeItem('login');
        window.location.href="/"    
    }
    return (
        <>
            <IconContext.Provider value={{ color: 'white' }}>
                <nav className={sidebar ? 'nav-menu active' : 'nav-menu'}>
                    <ul className='nav-menu-items' onClick={showSidebar}>
                        <li className='nav-text'>
                            <Link to='#'>
                                {/* <FcIcons.FcMenu /> */}
                                <Image src={bsim} width="22" height="22"/>
                                <span className='span'>Simas Kredit</span>
                            </Link>
                        </li>
                        {Menu.map((item, index) => {
                            return (
                                <li key={index} className={item.cName}>
                                    <Link to={item.path}>
                                        {item.icon}
                                        <span className='span'>{item.title}</span>
                                    </Link>
                                </li>
                            );
                        })}
                    </ul>
                    <div>
                        <li className='nav-text position-absolute bottom-0'>
                            <Link onClick={() => logout()}>
                                <FcIcons.FcImport />
                                <span className='span'>Log Out</span>
                            </Link>
                        </li>
                    </div>
                </nav>
            </IconContext.Provider>
        </>
    )
}

export default Sidebars 