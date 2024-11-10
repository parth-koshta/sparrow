# Use the official hibiken/asynqmon image as base
FROM hibiken/asynqmon:latest

# Expose port 8081 instead of the default 8080
EXPOSE 8081

ENV PORT=8081

# Command to start the AsynqMon UI server
CMD ["./asynqmon", "--port", "8081", "--redis-addr", "redis:6379"]
