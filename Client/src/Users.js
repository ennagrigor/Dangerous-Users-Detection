import React, {Component} from "react";
import {UncontrolledCarousel} from 'reactstrap';
import axios from "axios";
import "./Users.css"

class Users extends Component {
    constructor(props) {
        super(props);
        this.state = {
            users: []
        }
    }

    componentDidMount() {
        this.getDangerousUsers()
    }

    getDangerousUsers = () => {
        axios.get("http://localhost:8080/dangerous-users")
            .then(response => {
                const users = response?.data?.map((user, index) => ({
                    src: user?.profileImage,
                    altText: user?.username,
                    header: user?.username,
                    caption: `Score: ${user?.score}`,
                    key: index
                }));
                this.setState({
                    users
                })
            })
    }

    render() {
        const {users} = this.state
        if (!users) {
            return
        }
        const header = <h2>TOP DANGEROUS USERS</h2>;

        return (
            <div className="users">
                {header}
                <div className="dangerous-users">
                    <UncontrolledCarousel items={users}/>
                </div>
            </div>
        );
    }
}

export default Users
