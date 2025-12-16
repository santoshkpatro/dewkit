import asyncio
import logging

import tornado.web
import tornado.websocket
import tornado.log

from django.core.management.base import BaseCommand


class RelayWebSocket(tornado.websocket.WebSocketHandler):
    def open(self):
        logging.info(
            "WebSocket opened from %s",
            self.request.remote_ip,
        )

    def on_message(self, message):
        logging.debug("Received message: %s", message)
        # Echo back (replace with your relay logic)
        self.write_message(f"Echo: {message}")

    def on_close(self):
        logging.info("WebSocket closed")

    def check_origin(self, origin):
        # WARNING: lock this down in production
        return True


def make_app():
    return tornado.web.Application(
        [
            (r"/ws", RelayWebSocket),
        ],
        debug=False,
    )


async def main(port: int):
    app = make_app()
    app.listen(port, address="0.0.0.0")

    logging.info("Tornado WebSocket server listening on port %s", port)

    # Keep the server running forever
    await asyncio.Event().wait()


class Command(BaseCommand):
    help = "Start Tornado WebSocket relay server"

    def add_arguments(self, parser):
        parser.add_argument(
            "--port",
            type=int,
            default=9000,
            help="Port to run WebSocket server on",
        )
        parser.add_argument(
            "--log-level",
            default="INFO",
            choices=["DEBUG", "INFO", "WARNING", "ERROR", "CRITICAL"],
            help="Logging level",
        )

    def handle(self, *args, **options):
        self.setup_logging(options["log_level"])
        asyncio.run(main(options["port"]))

    def setup_logging(self, log_level: str):
        logging.basicConfig(
            level=getattr(logging, log_level),
            format="%(asctime)s [%(levelname)s] %(name)s: %(message)s",
        )

        # Make Tornado logs go through stdlib logging
        tornado.log.enable_pretty_logging()
