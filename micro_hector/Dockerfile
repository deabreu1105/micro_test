FROM python:3.10
WORKDIR /app
COPY requirements.txt .
RUN pip install --trusted-host pypi.python.org -r requirements.txt
RUN pip install Flask
RUN pip install flask-cors
COPY . .
ENV PORT 8080
# Run app.py when the container launches
CMD ["python", "main.py"]
