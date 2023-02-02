package com.stock.server.config.db.migrations.mariadb;

import org.flywaydb.core.Flyway;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.context.annotation.Configuration;

@Configuration
public class FlywayConfig {
    public FlywayConfig(@Qualifier("mariadbFlyway") Flyway mariadbFlyway) {
        mariadbFlyway.migrate();
    }
}
