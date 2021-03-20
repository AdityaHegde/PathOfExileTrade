import { createLogger, Logger, transports, format } from "winston";

export class InternalLogger {
  private logger: Logger;
  private label: string;

  constructor(label: string) {
    this.logger = createLogger({
      level: "debug",
      transports: [new transports.Console({
        format: format.simple(),
      })],
    });
    this.label = label;
  }

  public debug(message: string) {
    this.logger.debug(message);
  }

  public info(message: string) {
    this.logger.info(message);
  }

  public warn(message: string) {
    this.logger.warn(message);
  }
}
