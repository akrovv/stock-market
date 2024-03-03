import React , { useState , useEffect } from 'react';

import { Link , NavLink } from 'react-router-dom';
import menus from '../../pages/menu';
import { Dropdown } from 'react-bootstrap';

import './styles.scss';
import avt from '../../assets/images/avt/avt-01.jpg';
import DarkMode from './DarkMode';
import { useAuth } from '../../AuthContext';
import { useNavigate } from "react-router-dom";



const Header = () => {
    const [scroll, setScroll] = useState(false);
        useEffect(() => {
        window.addEventListener("scroll", () => {
            setScroll(window.scrollY > 300);
        });
        return () => {
            setScroll({});
        }
    }, []);

    const [menuActive, setMenuActive] = useState(null);

    const handleMenuActive = () => {
        setMenuActive(!menuActive);
      };

    const [activeIndex, setActiveIndex] = useState(null);
    const handleDropdown = index => {
        setActiveIndex(index); 
    };

    const { logout } = useAuth();
    const { isLoggedIn } = useAuth();
    const navigate = useNavigate();

    const handleLogout = async () => {
        localStorage.removeItem('token');
        logout();
        
        try {
            const response = await fetch("/api/user/logout", {
                method: "GET",
            });
    
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }
    
            navigate("/");
        } catch (error) {
            console.error(error);
        }
    };

    return (
        <header id="header_main" className={`header ${scroll ? 'is-fixed' : ''}`}>
            <div className="container-fluid">
                <div className="row">
                <div className="col-12">
                    <div className="header__body d-flex justify-content-between">
                    <div className="header__left">
                        <div className="left__main">
                            <nav id="main-nav" className={`main-nav ${menuActive ? 'active' : ''}`}>
                                <ul id="menu-primary-menu" className="menu">
                                {
                                    menus.map((data,idx) => (
                                        <li key={idx} onClick={()=> handleDropdown(idx)} className={`menu-item ${data.namesub ? 'menu-item-has-children' : ''} ${activeIndex === idx ? 'active' : ''}`} 
                                        
                                        >
                                            <Link to={data.links}>{data.name}</Link>
                                            {
                                                data.namesub &&
                                                <ul className="sub-menu">
                                                    {
                                                        data.namesub.map((submenu) => (
                                                            <li key={submenu.id} className="menu-item"><NavLink to={submenu.links}>{submenu.sub}</NavLink></li>
                                                        ))
                                                    }
                                                </ul>
                                            }
                                            
                                        </li>
                                    ))
                                }
                                </ul>
                            </nav>
                            
                 
                        </div>
                    </div>

                    <div className="header__right">
                        <DarkMode />
                        
                            <div className={`mobile-button ${menuActive ? 'active' : ''}`} onClick={handleMenuActive}><span></span></div>
                            {isLoggedIn ? (
                            <div className="wallet">
                            <Link to="/wallet"> Wallet </Link>
                            </div>
                            ) : (
                                <div className="user">
                                    <Link to="/login"><span className='me-4'>Login</span></Link>
                                    <Link to="/register"><span>Register</span></Link>
                                </div>
                            )}
                            {isLoggedIn && (
                            <Dropdown className='user'>
                                    <Dropdown.Toggle >
                                        <img src={avt} alt="Rockie" />  
                                    </Dropdown.Toggle>

                                    <Dropdown.Menu>
                                    <Dropdown.Item href="#">
                                        <Link className="dropdown-item" to="#"><i className="bx bx-user font-size-16 align-middle me-1"></i>
                                        <span>Profile</span></Link>
                                    </Dropdown.Item>
                                    <Dropdown.Item href="#">
                                        <Link className="dropdown-item" to="#"><i
                                            className="bx bx-wallet font-size-16 align-middle me-1"
                                        ></i>
                                        <span>My Wallet</span></Link>
                                    </Dropdown.Item>
                                    <Dropdown.Item href="#">
                                        <Link className="dropdown-item text-danger" to="/login"
                                        ><i
                                            className="bx bx-power-off font-size-16 align-middle me-1 text-danger"
                                        ></i>
                                        <span onClick={handleLogout}>Logout</span></Link>
                                    </Dropdown.Item>
                                    
                                    </Dropdown.Menu>
                                </Dropdown>
                        )}
                    </div>
                    </div>
                </div>
                </div>
            </div>
        </header>
       
    );
}

export default Header;