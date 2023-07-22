import pytest
from main import app

def test_status_code_200():
    with app.test_client() as client:
        response = client.get('/saludo')
        assert response.status_code == 200