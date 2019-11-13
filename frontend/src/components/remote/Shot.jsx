import { Button } from 'react-bootstrap';
import React from 'react';
import styled from 'styled-components';

class Shot extends React.Component {
    render() {
        return (
            <ColBase className="col-sm">
                <div className="col">
                    <h1 className="display-4">{this.props.status}</h1>
                </div>
                <div className="col">
                    <p>{this.props.name}</p>
                    <Button variant="primary">
                        {this.props.icon !== undefined ? this.props.icon : this.props.contents}
                    </Button>
                </div>
            </ColBase>
        )
    }
}

const ColBase = styled.div`
    margin: 12px;
`;

export default Shot;
