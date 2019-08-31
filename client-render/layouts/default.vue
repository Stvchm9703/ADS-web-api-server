<template lang="pug">
div
  nav.navbar.header.has-shadow.is-primary(role='navigation' aria-label='main navigation')
    .navbar-brand
      a.navbar-item(href='/')
        h2.title.has-text-white SCOPE ADS
      .navbar-burger
        span
        span
        span
  section.main-content.columns
    aside.column.is-2.section
      p.menu-label.is-hidden-touch
        | General
      ul.menu-list
        li
          nuxt-link(to='/' exact-active-class='is-active')
            b-icon(icon='home')
            | Home
      p.menu-label.is-hidden-touch
        | Demo Work
      ul.menu-list
        li(v-for='(item, key) of items' :key='key')
          nuxt-link(:to='item.to' exact-active-class='is-active')
            b-icon(:icon='item.icon')
            | {{ item.title }}
      p.menu-label.is-hidden-touch
        | Document for this Project
      ul.menu-list
        li
          nuxt-link(
            to='/doc' 
            exact-active-class='is-active')
            b-icon(icon='home')
            | Info and Summary 
        li
          nuxt-link(
            to='/doc#api' 
            exact-active-class='is-active')
            | API
        
          ul
            nuxt-link(
              to='/doc#api/course' 
              exact-active-class='is-active')
              | Course
          ul
            nuxt-link(
              to='/doc#api/department' 
              exact-active-class='is-active')
              | Department
          ul
            nuxt-link(
              to='/doc#api/enroll' 
              exact-active-class='is-active')
              | Enroll
          ul
            nuxt-link(
              to='/doc#api/offer' 
              exact-active-class='is-active')
              | Offer
          ul
            nuxt-link(
              to='/doc#api/student' 
              exact-active-class='is-active')
              | Student
          ul
            nuxt-link(
              to='/doc#api/vstudentenroll' 
              exact-active-class='is-active')
              | Student Enrolled Course (Read Only)
          ul
            nuxt-link(
              to='/doc#api/vcourse' 
              exact-active-class='is-active')
              | Course Offer Detail (Read Only)

    .container.column.is-10
      nuxt

</template>

<script>
export default {
  data: () => ({
    items: [
      {
        title: "Department",
        icon: "office-building",
        to: {
          name: "dept"
        }
      },
      {
        title: "Student",
        icon: "account",
        to: {
          name: "student"
        }
      }
    ],
    docList: []
  })
};
</script>
