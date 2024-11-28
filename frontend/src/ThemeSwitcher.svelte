<!-- <script> -->
<!--   import { onMount } from "svelte"; -->
<!---->
<!--   let currentTheme = "macchiato"; -->
<!---->
<!--   function toggleTheme() { -->
<!--     currentTheme = currentTheme === "macchiato" ? "latte" : "macchiato"; -->
<!--     document.getElementById("app").setAttribute("data-theme", currentTheme); -->
<!--   } -->
<!---->
<!--   onMount(() => { -->
<!--     currentTheme = document.getElementById("app").getAttribute("data-theme"); -->
<!--   }); -->
<!-- </script> -->
<!---->
<!-- <button on:click={toggleTheme} class="btn btn-primary"> -->
<!--   Switch to {currentTheme === "macchiato" ? "Latte" : "Macchiato"} -->
<!-- </button> -->
<!---->
<script>
  import { onMount } from "svelte";

  let currentTheme;

  function toggleTheme() {
    currentTheme = currentTheme === "macchiato" ? "latte" : "macchiato";
    document.getElementById("app").setAttribute("data-theme", currentTheme);
    localStorage.setItem("theme", currentTheme);
  }

  onMount(() => {
    currentTheme = localStorage.getItem("theme") || "macchiato";
    document.getElementById("app").setAttribute("data-theme", currentTheme);
  });

  $: isChecked = currentTheme === "macchiato";
</script>

<label class="swap swap-rotate">
  <input
    type="checkbox"
    class="toggle theme-controller"
    bind:checked={isChecked}
    on:change={toggleTheme}
  />
  <span class="swap-on">🌞</span>
  <span class="swap-off">🌙</span>
</label>
