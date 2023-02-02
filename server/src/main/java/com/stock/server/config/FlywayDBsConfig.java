package com.stock.server.config;

import org.flywaydb.core.Flyway;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class FlywayDBsConfig {
//    @Bean
//    @Qualifier("mariadbFlyway")
//    public Flyway mariadbFlyway(
//            @Value("${spring.flyway.mariadb.url}") String url,
//            @Value("${spring.flyway.mariadb.user}") String user,
//            @Value("${spring.flyway.mariadb.password}") String password,
//            @Value("${spring.flyway.mariadb.default-schema}") String defaultSchema,
//            @Value("${spring.flyway.mariadb.locations}") String location
//    ) {
//        return Flyway.configure()
//                .dataSource(url, user, password)
//                .locations(location)
//                .defaultSchema(defaultSchema)
//                .load();
//    }

    @Bean
    @Qualifier("cassandraFlyway")
    public Flyway cassandraFlyway(
            @Value("${spring.flyway.cassandra.url}") String url,
            @Value("${spring.flyway.cassandra.user}") String user,
            @Value("${spring.flyway.cassandra.password}") String password,
//            @Value("${spring.flyway.cassandra.default-schema}") String defaultSchema,
            @Value("${spring.flyway.cassandra.locations}") String location
    ) {
        return Flyway.configure()
                .dataSource(url, user, password)
                .locations(location)
//                .defaultSchema(defaultSchema)
                .load();
    }
}
