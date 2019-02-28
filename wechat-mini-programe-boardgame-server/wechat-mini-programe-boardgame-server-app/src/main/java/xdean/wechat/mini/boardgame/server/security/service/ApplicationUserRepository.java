package xdean.wechat.mini.boardgame.server.security.service;

import org.springframework.data.jpa.repository.JpaRepository;

import xdean.wechat.mini.boardgame.server.security.model.ApplicationUser;

public interface ApplicationUserRepository extends JpaRepository<ApplicationUser, Long> {
  ApplicationUser findByUsername(String username);
}