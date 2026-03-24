import logging
from typing import Optional
from pathlib import Path

class CoreEngine:
    def __init__(self, config_file: Optional[str] = None):
        self.config_file = config_file
        self.logger = self._setup_logger()

    def _setup_logger(self) -> logging.Logger:
        logger = logging.getLogger("core-engine")
        logger.setLevel(logging.INFO)
        handler = logging.StreamHandler()
        formatter = logging.Formatter("%(asctime)s - %(name)s - %(levelname)s - %(message)s")
        handler.setFormatter(formatter)
        logger.addHandler(handler)
        return logger

    def load_config(self) -> dict:
        if not self.config_file:
            self.logger.warning("No config file provided. Using default settings.")
            return {"default": True}

        config_path = Path(self.config_file)
        if not config_path.exists():
            self.logger.error(f"Config file {self.config_file} not found.")
            raise FileNotFoundError(f"Config file {self.config_file} not found.")

        try:
            with open(config_path, "r") as file:
                # Simulate loading a config file (e.g., JSON, YAML)
                return {"loaded": True}
        except Exception as e:
            self.logger.error(f"Failed to load config file: {e}")
            raise

    def start(self) -> None:
        self.logger.info("Starting Core Engine...")
        config = self.load_config()
        self.logger.info(f"Engine started with config: {config}")

    def stop(self) -> None:
        self.logger.info("Stopping Core Engine...")

if __name__ == "__main__":
    engine = CoreEngine(config_file="config.json")
    engine.start()
    engine.stop()