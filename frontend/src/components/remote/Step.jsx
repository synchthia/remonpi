import { Button, ButtonGroup } from 'react-bootstrap';
import React from 'react';
import styled from 'styled-components';

class Step extends React.Component {
    render() {
        return (
            <ColBase className="col-sm">
                <div className="col">
                    <h1 className="display-4">{this.props.status}</h1>
                </div>
                <div className="col">
                    <p>{this.props.name}</p>
                    <ButtonGroup>
                        {this.props.contents.map((e) => {
                            return (
                                <Button
                                    key={e}
                                    variant={e === this.props.status ? "primary" : "secondary"}
                                    onClick={() => this.props.onClick(e)}
                                >
                                    {e}
                                </Button>
                            )
                        })}
                    </ButtonGroup>
                </div>
            </ColBase>
        )
    }
}

const ColBase = styled.div`
    margin: 12px;
`;

export default Step;
