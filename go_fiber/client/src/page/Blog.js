import axios from "axios";
import React, { useEffect, useState } from "react";
import { Col, Container, Row } from "react-bootstrap";
import { useParams } from "react-router-dom";

const Blog = () => {
    const params = useParams();

    console.log(params);
    const [apiData, setApiData] = useState(false);
    useEffect(() => {
        const fetchData = async () => {
    
          try {
            const apiUrl = process.env.REACT_APP_API_ROOT+ "/" +params.id;
            const response = await axios.get(apiUrl);
            
            if (response.status === 200) {
              if (response?.data.statusText === "OK") {
                setApiData(response?.data?.record);
              }
            }
          } catch (error) {
            console.log(error.response);
          }
    
        };
    
        fetchData();
        return () => {};
      }, [params.id]);


    console.log(apiData);


    return (
        <Container>
            <Row>
                <Col xs="12">
                    <h1>{apiData.title}</h1>
                </Col>
                <Col xs="12">
                    <p>{apiData.post}</p>
                </Col>
            </Row>
        </Container>
    )
}

export default Blog